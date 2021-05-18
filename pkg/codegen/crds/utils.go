package crds

import "github.com/rancher/wrangler/pkg/controller-gen/args"

func GetGroups(customResourceDefinitions []CustomResourceDefinition) map[string]args.Group {
	groups := make(map[string]args.Group)
	for _, crd := range customResourceDefinitions {
		var group args.Group
		var exists bool
		if group, exists = groups[crd.APIGroup]; !exists {
			group = args.Group{GenerateTypes: true}
		}
		if group.Types == nil {
			group.Types = make([]interface{}, 0)
		}
		group.Types = append(group.Types, crd.Type)
		groups[crd.APIGroup] = group
	}
	return groups
}

func WriteCRDs(customResourceDefinitions []CustomResourceDefinition, toDir string) error {
	for _, crd := range customResourceDefinitions {
		err := crd.WriteCRD(toDir)
		if err != nil {
			return err
		}
	}
	return nil
}
