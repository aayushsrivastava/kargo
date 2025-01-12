package api

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"connectrpc.com/connect"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	kargoapi "github.com/akuity/kargo/api/v1alpha1"
	svcv1alpha1 "github.com/akuity/kargo/pkg/api/service/v1alpha1"
)

func (s *server) ListProjectSecrets(
	ctx context.Context,
	req *connect.Request[svcv1alpha1.ListProjectSecretsRequest],
) (*connect.Response[svcv1alpha1.ListProjectSecretsResponse], error) {
	// Check if secret management is enabled
	if !s.cfg.SecretManagementEnabled {
		return nil, connect.NewError(connect.CodeUnimplemented, errSecretManagementDisabled)
	}

	project := req.Msg.GetProject()
	if err := validateFieldNotEmpty("project", project); err != nil {
		return nil, err
	}

	if err := s.validateProjectExists(ctx, project); err != nil {
		return nil, err
	}

	var secretsList corev1.SecretList
	if err := s.client.List(
		ctx,
		&secretsList,
		client.InNamespace(req.Msg.GetProject()),
		client.MatchingLabels{
			kargoapi.ProjectSecretLabelKey: kargoapi.LabelTrueValue,
		},
	); err != nil {
		return nil, fmt.Errorf("list secrets: %w", err)
	}

	// Sort ascending by name
	slices.SortFunc(secretsList.Items, func(lhs, rhs corev1.Secret) int {
		return strings.Compare(lhs.Name, rhs.Name)
	})

	secrets := make([]*corev1.Secret, len(secretsList.Items))
	for i, secret := range secretsList.Items {
		secrets[i] = sanitizeProjectSecret(secret)
	}

	return connect.NewResponse(&svcv1alpha1.ListProjectSecretsResponse{
		Secrets: secrets,
	}), nil
}

// sanitizeProjectSecret returns a copy of the secret with all values in the
// stringData map redacted. All annotations are also redacted because AT LEAST
// "last-applied-configuration" is a known vector for leaking sensitive
// information and unknown configuration management tools may use other
// annotations in a manner similar to "last-applied-configuration". There is no
// concern over labels because the constraints on label values rule out use in a
// manner similar to that of the "last-applied-configuration" annotation.
func sanitizeProjectSecret(secret corev1.Secret) *corev1.Secret {
	s := secret.DeepCopy()
	s.StringData = make(map[string]string, len(s.Data))
	for k, v := range s.Annotations {
		switch k {
		case kargoapi.AnnotationKeyDescription:
			s.Annotations[k] = v
		default:
			s.Annotations[k] = redacted
		}
	}
	for k := range s.Data {
		s.StringData[k] = redacted
	}
	s.Data = nil
	return s
}