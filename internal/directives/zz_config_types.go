// Code generated by quicktype. DO NOT EDIT.

package directives

type CommonDefs interface{}

type ArgoCDUpdateConfig struct {
	Apps       []ArgoCDAppUpdate `json:"apps"`
	FromOrigin *AppFromOrigin    `json:"fromOrigin,omitempty"`
}

type ArgoCDAppUpdate struct {
	FromOrigin *AppFromOrigin `json:"fromOrigin,omitempty"`
	// Specifies the name of an Argo CD Application resource to be updated.
	Name string `json:"name"`
	// Specifies the namespace of an Argo CD Application resource to be updated. If left
	// unspecified, the namespace will be the controller's configured default.
	Namespace string `json:"namespace,omitempty"`
	// Describes updates to be applied to various sources of an Argo CD Application resource.
	Sources []ArgoCDAppSourceUpdate `json:"sources,omitempty"`
}

type AppFromOrigin struct {
	// The kind of origin. Currently only 'Warehouse' is supported. Required.
	Kind Kind `json:"kind"`
	// The name of the origin. Required.
	Name string `json:"name"`
}

type ArgoCDAppSourceUpdate struct {
	// If applicable, identifies a specific chart within the Helm chart repository specified by
	// the 'repoURL' field. When the source to be updated references a Helm chart repository,
	// the values of the 'repoURL' and 'chart' fields should exactly match the values of the
	// same fields in the source. i.e. Do not match the values of these two fields to your
	// Warehouse; match them to the Application source you wish to update.
	Chart      string                       `json:"chart,omitempty"`
	FromOrigin *AppFromOrigin               `json:"fromOrigin,omitempty"`
	Helm       *ArgoCDHelmParameterUpdates  `json:"helm,omitempty"`
	Kustomize  *ArgoCDKustomizeImageUpdates `json:"kustomize,omitempty"`
	// With possible help from the 'chart' field, identifies which of an Argo CD Application's
	// sources is to be updated. When the source to be updated references a Helm chart
	// repository, the values of the 'repoURL' and 'chart' fields should exactly match the
	// values of the same fields in the source. i.e. Do not match the values of these two fields
	// to your Warehouse; match them to the Application source you wish to update.
	RepoURL string `json:"repoURL"`
	// Indicates whether the source should be updated such that its TargetRevision field points
	// at the most recently git commit (if 'repoURL' references a Git repository) or chart
	// version (if 'repoURL' references a chart repository).
	UpdateTargetRevision bool `json:"updateTargetRevision,omitempty"`
}

// Describes updates to an Argo CD Application source's Helm parameters.
type ArgoCDHelmParameterUpdates struct {
	FromOrigin *AppFromOrigin          `json:"fromOrigin,omitempty"`
	Images     []ArgoCDHelmImageUpdate `json:"images"`
}

// Describes how to update a Helm parameter to reference a specific version of a container
// image.
type ArgoCDHelmImageUpdate struct {
	FromOrigin *AppFromOrigin `json:"fromOrigin,omitempty"`
	// Specifies a key within an Argo CD Application source's Helm parameters that is to be
	// updated.
	Key string `json:"key"`
	// The URL of a container image repository.
	RepoURL string `json:"repoURL"`
	// Specifies a new value for the setting within an Argo CD Application source's Helm
	// parameters identified by the 'key' field.
	Value Value `json:"value"`
}

// Describes updates to an Argo CD Application source's Kustomize images.
type ArgoCDKustomizeImageUpdates struct {
	FromOrigin *AppFromOrigin               `json:"fromOrigin,omitempty"`
	Images     []ArgoCDKustomizeImageUpdate `json:"images,omitempty"`
}

// Describes how to update a Kustomize image to reference a specific version of a container
// image.
type ArgoCDKustomizeImageUpdate struct {
	FromOrigin *AppFromOrigin `json:"fromOrigin,omitempty"`
	// Specifies a container image name override.
	NewName string `json:"newName,omitempty"`
	// The URL of a container image repository.
	RepoURL string `json:"repoURL"`
	// Specifies whether the image's digest should be used instead of its tag.
	UseDigest bool `json:"useDigest,omitempty"`
}

type CopyConfig struct {
	// InPath is the path to the file or directory to copy.
	InPath string `json:"inPath"`
	// OutPath is the path to the destination file or directory.
	OutPath string `json:"outPath"`
}

type GitCloneConfig struct {
	// The commits, branches, or tags to check out from the repository and the paths where they
	// should be checked out. At least one must be specified.
	Checkout []Checkout `json:"checkout"`
	// Indicates whether to skip TLS verification when cloning the repository. Default is false.
	InsecureSkipTLSVerify bool `json:"insecureSkipTLSVerify,omitempty"`
	// The URL of a remote Git repository to clone. Required.
	RepoURL string `json:"repoURL"`
}

type Checkout struct {
	// The branch to checkout. Mutually exclusive with 'tag' and 'fromFreight=true'. If none of
	// these is specified, the default branch is checked out.
	Branch string `json:"branch,omitempty"`
	// Indicates whether a new, empty orphan branch should be created if the branch does not
	// already exist. Default is false.
	Create bool `json:"create,omitempty"`
	// Indicates whether the ID of a commit to check out may be obtained from Freight. A value
	// of 'true' is mutually exclusive with 'branch' and 'tag'. If none of these is specified,
	// the default branch is checked out.
	FromFreight bool                `json:"fromFreight,omitempty"`
	FromOrigin  *CheckoutFromOrigin `json:"fromOrigin,omitempty"`
	// The path where the repository should be checked out.
	Path string `json:"path"`
	// The tag to checkout. Mutually exclusive with 'branch' and 'fromFreight=true'. If none of
	// these is specified, the default branch is checked out.
	Tag string `json:"tag,omitempty"`
}

type CheckoutFromOrigin struct {
	// The kind of origin. Currently only 'Warehouse' is supported. Required.
	Kind Kind `json:"kind"`
	// The name of the origin. Required.
	Name string `json:"name"`
}

type GitCommitConfig struct {
	// The author of the commit.
	Author *Author `json:"author,omitempty"`
	// The commit message. Mutually exclusive with 'messageFrom'.
	Message string `json:"message,omitempty"`
	// TODO
	MessageFrom []string `json:"messageFrom,omitempty"`
	// The path to a working directory of a local repository.
	Path string `json:"path"`
}

// The author of the commit.
type Author struct {
	// The email of the author.
	Email string `json:"email,omitempty"`
	// The name of the author.
	Name string `json:"name,omitempty"`
}

type GitOpenPRConfig struct {
	// Indicates whether a new, empty orphan branch should be created and pushed to the remote
	// if the target branch does not already exist there. Default is false.
	CreateTargetBranch bool `json:"createTargetBranch,omitempty"`
	// Indicates whether to skip TLS verification when cloning the repository. Default is false.
	InsecureSkipTLSVerify bool `json:"insecureSkipTLSVerify,omitempty"`
	// The name of the Git provider to use. Currently only 'github' and 'gitlab' are supported.
	// Kargo will try to infer the provider if it is not explicitly specified.
	Provider *Provider `json:"provider,omitempty"`
	// The URL of a remote Git repository to clone.
	RepoURL string `json:"repoURL"`
	// The branch containing the changes to be merged. This branch must already exist and be up
	// to date on the remote.
	SourceBranch string `json:"sourceBranch,omitempty"`
	// References the 'branch' output from a previous step. This step will use that value as the
	// source branch.
	SourceBranchFrom string `json:"sourceBranchFrom,omitempty"`
	// The branch to which the changes should be merged. This branch must already exist and be
	// up to date on the remote.
	TargetBranch string `json:"targetBranch"`
}

type GitOverwriteConfig struct {
	// A path to a directory from which to copy all contents, excluding the .git/ directory, if
	// one exists.
	InPath string `json:"inPath"`
	// A path to a git working tree which will be cleared of all existing content before
	// receiving a copy of all content specified by inPath.
	OutPath string `json:"outPath"`
}

type GitPushConfig struct {
	// Indicates whether to push to a new remote branch. A value of 'true' is mutually exclusive
	// with 'targetBranch'. If neither of these is provided, the target branch will be the
	// currently checked out branch.
	GenerateTargetBranch bool `json:"generateTargetBranch,omitempty"`
	// The path to a working directory of a local repository.
	Path string `json:"path"`
	// The target branch to push to. Mutually exclusive with 'generateTargetBranch=true'. If
	// neither of these is provided, the target branch will be the currently checked out branch.
	TargetBranch string `json:"targetBranch,omitempty"`
}

type GitWaitForPRConfig struct {
	// Indicates whether to skip TLS verification when cloning the repository. Default is false.
	InsecureSkipTLSVerify bool `json:"insecureSkipTLSVerify,omitempty"`
	// The number of the pull request to wait for.
	PRNumber int64 `json:"prNumber,omitempty"`
	// References a previous open step by alias and will use the PR number opened by that step.
	PRNumberFromOpen string `json:"prNumberFromOpen,omitempty"`
	// The name of the Git provider to use. Currently only 'github' and 'gitlab' are supported.
	// Kargo will try to infer the provider if it is not explicitly specified.
	Provider *Provider `json:"provider,omitempty"`
	// The URL of a remote Git repository to clone.
	RepoURL string `json:"repoURL"`
}

type HelmTemplateConfig struct {
	// APIVersions allows a manual set of supported API Versions to be passed when rendering the
	// manifests.
	APIVersions []string `json:"apiVersions,omitempty"`
	// Whether to include CRDs in the rendered manifests.
	IncludeCRDs bool `json:"includeCRDs,omitempty"`
	// KubeVersion allows for passing a specific Kubernetes version to use when rendering the
	// manifests.
	KubeVersion string `json:"kubeVersion,omitempty"`
	// Namespace to use for the rendered manifests.
	Namespace string `json:"namespace,omitempty"`
	// OutPath to write the rendered manifests to.
	OutPath string `json:"outPath"`
	// Path at which the Helm chart can be found.
	Path string `json:"path"`
	// ReleaseName to use for the rendered manifests.
	ReleaseName string `json:"releaseName,omitempty"`
	// ValuesFiles to use for rendering the Helm chart.
	ValuesFiles []string `json:"valuesFiles,omitempty"`
}

type HelmUpdateChartConfig struct {
	// A list of chart dependencies which should receive updates.
	Charts []Chart `json:"charts"`
	// The path at which the umbrella chart with the dependency can be found.
	Path string `json:"path"`
}

type Chart struct {
	FromOrigin *ChartFromOrigin `json:"fromOrigin,omitempty"`
	// The name of the subchart, as defined in `Chart.yaml`.
	Name string `json:"name"`
	// The repository of the subchart, as defined in `Chart.yaml`. It also supports OCI charts
	// using `oci://`.
	Repository string `json:"repository"`
}

type ChartFromOrigin struct {
	// The kind of origin. Currently only 'Warehouse' is supported. Required.
	Kind Kind `json:"kind"`
	// The name of the origin. Required.
	Name string `json:"name"`
}

type HelmUpdateImageConfig struct {
	// A list of images which should receive updates.
	Images []HelmUpdateImageConfigImage `json:"images"`
	// The path at which the Helm values file can be found.
	Path string `json:"path"`
}

type HelmUpdateImageConfigImage struct {
	FromOrigin *ChartFromOrigin `json:"fromOrigin,omitempty"`
	// The container image (without tag) at which the update is targeted.
	Image string `json:"image"`
	// The key in the Helm values file of which the value needs to be updated. For nested
	// values, it takes a YAML dot notation path.
	Key string `json:"key"`
	// Specifies the new value for the specified key in the Helm values file.
	Value Value `json:"value"`
}

type KustomizeBuildConfig struct {
	// OutPath is the file path to write the built manifests to.
	OutPath string `json:"outPath"`
	// Path to the directory containing the Kustomization file.
	Path string `json:"path"`
}

type KustomizeSetImageConfig struct {
	// Images is a list of container images to set or update in the Kustomization file.
	Images []KustomizeSetImageConfigImage `json:"images"`
	// Path to the directory containing the Kustomization file.
	Path string `json:"path"`
}

type KustomizeSetImageConfigImage struct {
	FromOrigin *ChartFromOrigin `json:"fromOrigin,omitempty"`
	// Image name of the repository from which to pick the version. This is the image name Kargo
	// is subscribed to, and produces Freight for.
	Image string `json:"image"`
	// Name of the image (as defined in the Kustomization file).
	Name string `json:"name,omitempty"`
	// NewName for the image. This can be used to rename the container image name in the
	// manifests.
	NewName string `json:"newName,omitempty"`
	// UseDigest specifies whether to use the digest of the image instead of the tag.
	UseDigest bool `json:"useDigest,omitempty"`
}

// The kind of origin. Currently only 'Warehouse' is supported. Required.
type Kind string

const (
	Warehouse Kind = "Warehouse"
)

// Specifies a new value for the setting within an Argo CD Application source's Helm
// parameters identified by the 'key' field.
//
// Specifies the new value for the specified key in the Helm values file.
type Value string

const (
	Digest         Value = "Digest"
	ImageAndDigest Value = "ImageAndDigest"
	ImageAndTag    Value = "ImageAndTag"
	Tag            Value = "Tag"
)

// The name of the Git provider to use. Currently only 'github' and 'gitlab' are supported.
// Kargo will try to infer the provider if it is not explicitly specified.
type Provider string

const (
	Github Provider = "github"
	Gitlab Provider = "gitlab"
)
