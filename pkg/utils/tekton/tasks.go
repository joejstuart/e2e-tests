package tekton

import (
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// const (
// 	kanikoTaskName = "kaniko-chains"
// )

func KanikoTaskRun(namespace string) *v1beta1.TaskRun {
	return &v1beta1.TaskRun{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "kaniko-taskrun",
			Namespace:    namespace,
		},
		Spec: v1beta1.TaskRunSpec{
			Params: []v1beta1.Param{
				{
					Name: "IMAGE",
					Value: v1beta1.ArrayOrString{
						Type:      v1beta1.ParamTypeString,
						StringVal: "image-registry.openshift-image-registry.svc:5000/tekton-chains/kaniko-chains",
					},
				},
			},
			TaskRef: &v1beta1.TaskRef{
				Kind:   v1beta1.NamespacedTaskKind,
				Name:   "kaniko-chains",
				Bundle: "quay.io/jstuart/appstudio-tasks:latest-1",
			},
			Workspaces: []v1beta1.WorkspaceBinding{
				{
					Name:     "source",
					EmptyDir: &corev1.EmptyDirVolumeSource{},
				},
			},
		},
	}
}

// func KanikoTask(namespace, destinationImage string) *v1beta1.Task {
// 	ref, err := name.ParseReference(destinationImage)
// 	if err != nil {
// 		fmt.Printf("unable to parse image name: %v", err)
// 	}
// 	return &v1beta1.Task{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      kanikoTaskName,
// 			Namespace: namespace,
// 		},
// 		Spec: v1beta1.TaskSpec{
// 			Results: []v1beta1.TaskResult{
// 				{Name: "IMAGE_URL"},
// 				{Name: "IMAGE_DIGEST"},
// 			},
// 			Workspaces: []v1beta1.WorkspaceDeclaration{
// 				{
// 					Name: "source",
// 				},
// 				{
// 					Name:      "dockerconfig",
// 					Optional:  true,
// 					MountPath: "/kaniko/.docker",
// 				},
// 			},
// 			Steps: []v1beta1.Step{
// 				{
// 					Container: v1.Container{
// 						Name:       "add-dockerfile",
// 						Image:      "bash:latest",
// 						WorkingDir: "$(workspaces.source.path)",
// 					},
// 					Script: "#!/usr/bin/env bash\necho \"FROM alpine@sha256:69e70a79f2d41ab5d637de98c1e0b055206ba40a8145e7bddb55ccc04e13cf8f\" | tee ./Dockerfile",
// 				}, {
// 					Container: v1.Container{
// 						Name:  "build-and-push",
// 						Image: "gcr.io/kaniko-project/executor:v1.5.1@sha256:c6166717f7fe0b7da44908c986137ecfeab21f31ec3992f6e128fff8a94be8a5",
// 						Args: []string{
// 							"--dockerfile=./Dockerfile",
// 							fmt.Sprintf("--destination=%s", destinationImage),
// 							"--context=$(workspaces.source.path)/./",
// 							"--digest-file=$(results.IMAGE_DIGEST.path)",
// 							// Need this to push the image to the insecure registry
// 							"--skip-tls-verify=true",
// 						},
// 						SecurityContext: &v1.SecurityContext{
// 							RunAsUser: new(int64),
// 						},
// 					},
// 				}, {
// 					Container: v1.Container{
// 						Name:  "write-url",
// 						Image: "bash:latest",
// 					},
// 					Script: fmt.Sprintf("#!/usr/bin/env bash\necho %s | tee $(results.IMAGE_URL.path)", ref.String()),
// 				},
// 			},
// 		},
// 	}
// }

func VerifyKanikoTaskRun(namespace string) *v1beta1.TaskRun {
	return &v1beta1.TaskRun{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "verify-kaniko-taskrun",
			Namespace:    namespace,
		},
		Spec: v1beta1.TaskRunSpec{
			Params: []v1beta1.Param{
				{
					Name: "TASK_LABEL",
					Value: v1beta1.ArrayOrString{
						Type:      v1beta1.ParamTypeString,
						StringVal: "kaniko-chains",
					},
				},
			},
			TaskRef: &v1beta1.TaskRef{
				Kind:   v1beta1.NamespacedTaskKind,
				Name:   "verify-attestation-signature",
				Bundle: "quay.io/jstuart/appstudio-tasks:latest-2",
			},
		},
	}
}
