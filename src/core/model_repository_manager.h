// Copyright (c) 2018-2019, NVIDIA CORPORATION. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
//  * Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
//  * Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
//  * Neither the name of NVIDIA CORPORATION nor the names of its
//    contributors may be used to endorse or promote products derived
//    from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS ``AS IS'' AND ANY
// EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
// PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL THE COPYRIGHT OWNER OR
// CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
// EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
// PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY
// OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
#pragma once

#include <mutex>
#include "src/core/backend.h"
#include "src/core/model_config.h"
#include "src/core/model_config.pb.h"
#include "src/core/status.h"
#include "src/servables/caffe2/netdef_bundle.h"
#include "src/servables/custom/custom_bundle.h"
#include "src/servables/ensemble/ensemble_bundle.h"
#include "src/servables/tensorflow/graphdef_bundle.h"
#include "src/servables/tensorflow/savedmodel_bundle.h"
#include "src/servables/tensorrt/plan_bundle.h"

namespace tensorflow { namespace serving {
class ModelConfig;
class ModelSpec;
class ServerCore;
class ServableHandle;
}}  // namespace tensorflow::serving
namespace tfs = tensorflow::serving;

namespace nvidia { namespace inferenceserver {

/// A singleton to manage the model repository active in the server. A
/// singleton is used because the servables have no connection to the
/// server itself but they need to have access to the configuration.
class ModelRepositoryManager {
 public:
  using VersionStateMap = std::map<int64_t, ModelReadyState>;
  using ModelMap = std::map<std::string, VersionStateMap>;

  enum ActionType {
    LOAD,
    UNLOAD
  };

  class BackendHandle {
   public:
    BackendHandle(const Platform& platform, const tfs::ModelSpec& spec, tfs::ServerCore* core);
    InferenceBackend* GetInferenceBackend() { return is_; }
   private:
    InferenceBackend* is_;
    tfs::ServableHandle<GraphDefBundle> graphdef_bundle_;
    tfs::ServableHandle<PlanBundle> plan_bundle_;
    tfs::ServableHandle<NetDefBundle> netdef_bundle_;
    tfs::ServableHandle<SavedModelBundle> saved_model_bundle_;
    tfs::ServableHandle<CustomBundle> custom_bundle_;
    tfs::ServableHandle<EnsembleBundle> ensemble_bundle_;
  };

  /// Create a manager for a repository.
  /// \param server_version The version of the inference server.
  /// \param status_manager The status manager that the model repository manager
  /// will update model configuration and state to.
  /// \param repositpory_path The file-system path of the repository.
  /// \param strict_model_config If false attempt to autofill missing required
  /// information in each model configuration.
  /// \param tf_gpu_memory_fraction The portion of GPU memory to be reserved
  /// for TensorFlow models.
  /// \param tf_allow_soft_placement If true instruct TensorFlow to use CPU
  /// implementation of an operation when a GPU implementation is not available
  /// \param repository_poll_secs Interval in seconds between each poll of
  /// the model repository to check for changes. [TODO] will be removed once
  /// ModelRepositoryManager is improved
  /// \return The error status.
  static Status Create(
      const std::string& server_version,
      const std::shared_ptr<ServerStatusManager>& status_manager,
      const std::string& repository_path,
      const bool strict_model_config, const float tf_gpu_memory_fraction,
      const bool tf_allow_soft_placement, const uint32_t repository_poll_secs);

  /// Poll the model repository to determine the new set of models and
  /// compare with the current set. And serve the new set of models based
  /// on their version policy.
  static Status PollAndUpdate();

  /// Load or unload a specified model.
  /// \parm model_name The name of the model to be loaded or unloaded
  /// \parm type The type action to be performed. If the action is LOAD and
  /// the model has been loaded, the model will be re-loaded.
  /// \param OnCompleteUpdate The callback function to be invoked once the
  /// action is completed.
  /// \return error status. Return "NOT_FOUND" if it tries to load
  /// a non-existing model or if it tries to unload a model that hasn't been
  /// loaded.
  static Status LoadUnloadModel(
      const std::string& model_name, ActionType type,
      std::function<void(Status)> OnCompleteUpdate);

  /// Unload all models.
  static Status UnloadAllModels();

  /// \param live_only only return the models that are live, at least one
  /// model version is available.
  /// \return the states of all versions of all model backends.
  static const ModelMap GetLiveBackendStates();

  /// \param model_name The model to get version states from.
  /// \return the states of all versions of the specified model backends.
  static const VersionStateMap GetVersionStates(const std::string& model_name);

  /// Obtain the specified backend handle.
  /// \param model_name The model name of the backend handle.
  /// \param model_version The model version of the backend handle.
  /// \param handle Return the backend handle object.
  /// \return error status.
  static Status GetBackendHandle(const std::string& model_name, const int64_t model_version, std::unique_ptr<BackendHandle>& handle);

  /// Poll the model repository to determine the new set of models and
  /// compare with the current set. Return the additions, deletions,
  /// and modifications that have occurred since the last Poll().
  /// \param added The names of the models added to the repository.
  /// \param deleted The names of the models removed from the repository.
  /// \param modified The names of the models remaining in the
  /// repository that have been changed.
  /// \param unmodified The names of the models remaining in the
  /// repository that have not changed.
  /// \return The error status.
  static Status Poll(
      std::set<std::string>* added, std::set<std::string>* deleted,
      std::set<std::string>* modified, std::set<std::string>* unmodified);

  /// Get the configuration for a named model.
  /// \param name The model name.
  /// \param model_config Returns the model configuration.
  /// \return OK if found, NOT_FOUND otherwise.
  static Status GetModelConfig(
      const std::string& name, ModelConfig* model_config);

  /// Get TFS-style configuration for a named model.
  /// \param name The model name.
  /// \param tfs_model_config Returns the TFS-style model configuration.
  /// \return OK if found, NOT_FOUND otherwise.
  static Status GetTFSModelConfig(
      const std::string& name, tfs::ModelConfig* tfs_model_config);

  /// Get the platform for a named model.
  /// \param name The model name.
  /// \param platform Returns the Platform.
  /// \return OK if found, NOT_FOUND otherwise.
  static Status GetModelPlatform(const std::string& name, Platform* platform);

 private:
  struct ModelInfo;

  // Map from model name to information about the model.
  using ModelInfoMap =
      std::unordered_map<std::string, std::unique_ptr<ModelInfo>>;

  ModelRepositoryManager(
      const std::shared_ptr<ServerStatusManager>& status_manager, 
      const std::string& repository_path,
      const PlatformConfigMap& platform_config_map, const bool autofill);
  ~ModelRepositoryManager() = default;

  static ModelRepositoryManager* singleton;

  const std::string repository_path_;
  const PlatformConfigMap platform_config_map_;
  const bool autofill_;

  std::mutex poll_mu_;
  std::mutex infos_mu_;
  ModelInfoMap infos_;

  std::unique_ptr<tensorflow::serving::ServerCore> core_;
  std::shared_ptr<ServerStatusManager> status_manager_;
};

}}  // namespace nvidia::inferenceserver
