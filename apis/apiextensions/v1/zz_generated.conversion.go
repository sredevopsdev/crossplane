// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package v1

import (
	v11 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	v1beta1 "github.com/crossplane/crossplane/apis/apiextensions/v1beta1"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type GeneratedRevisionSpecConverter struct{}

func (c *GeneratedRevisionSpecConverter) FromRevisionSpec(source v1beta1.CompositionRevisionSpec) CompositionSpec {
	var v1CompositionSpec CompositionSpec
	v1CompositionSpec.CompositeTypeRef = c.v1beta1TypeReferenceToV1TypeReference(source.CompositeTypeRef)
	v1PatchSetList := make([]PatchSet, len(source.PatchSets))
	for i := 0; i < len(source.PatchSets); i++ {
		v1PatchSetList[i] = c.v1beta1PatchSetToV1PatchSet(source.PatchSets[i])
	}
	v1CompositionSpec.PatchSets = v1PatchSetList
	var pV1EnvironmentConfiguration *EnvironmentConfiguration
	if source.Environment != nil {
		v1EnvironmentConfiguration := c.v1beta1EnvironmentConfigurationToV1EnvironmentConfiguration(*source.Environment)
		pV1EnvironmentConfiguration = &v1EnvironmentConfiguration
	}
	v1CompositionSpec.Environment = pV1EnvironmentConfiguration
	v1ComposedTemplateList := make([]ComposedTemplate, len(source.Resources))
	for j := 0; j < len(source.Resources); j++ {
		v1ComposedTemplateList[j] = c.v1beta1ComposedTemplateToV1ComposedTemplate(source.Resources[j])
	}
	v1CompositionSpec.Resources = v1ComposedTemplateList
	var pString *string
	if source.WriteConnectionSecretsToNamespace != nil {
		xstring := *source.WriteConnectionSecretsToNamespace
		pString = &xstring
	}
	v1CompositionSpec.WriteConnectionSecretsToNamespace = pString
	var pV1StoreConfigReference *StoreConfigReference
	if source.PublishConnectionDetailsWithStoreConfigRef != nil {
		v1StoreConfigReference := c.v1beta1StoreConfigReferenceToV1StoreConfigReference(*source.PublishConnectionDetailsWithStoreConfigRef)
		pV1StoreConfigReference = &v1StoreConfigReference
	}
	v1CompositionSpec.PublishConnectionDetailsWithStoreConfigRef = pV1StoreConfigReference
	return v1CompositionSpec
}
func (c *GeneratedRevisionSpecConverter) ToRevisionSpec(source CompositionSpec) v1beta1.CompositionRevisionSpec {
	var v1beta1CompositionRevisionSpec v1beta1.CompositionRevisionSpec
	v1beta1CompositionRevisionSpec.CompositeTypeRef = c.v1TypeReferenceToV1beta1TypeReference(source.CompositeTypeRef)
	v1beta1PatchSetList := make([]v1beta1.PatchSet, len(source.PatchSets))
	for i := 0; i < len(source.PatchSets); i++ {
		v1beta1PatchSetList[i] = c.v1PatchSetToV1beta1PatchSet(source.PatchSets[i])
	}
	v1beta1CompositionRevisionSpec.PatchSets = v1beta1PatchSetList
	var pV1beta1EnvironmentConfiguration *v1beta1.EnvironmentConfiguration
	if source.Environment != nil {
		v1beta1EnvironmentConfiguration := c.v1EnvironmentConfigurationToV1beta1EnvironmentConfiguration(*source.Environment)
		pV1beta1EnvironmentConfiguration = &v1beta1EnvironmentConfiguration
	}
	v1beta1CompositionRevisionSpec.Environment = pV1beta1EnvironmentConfiguration
	v1beta1ComposedTemplateList := make([]v1beta1.ComposedTemplate, len(source.Resources))
	for j := 0; j < len(source.Resources); j++ {
		v1beta1ComposedTemplateList[j] = c.v1ComposedTemplateToV1beta1ComposedTemplate(source.Resources[j])
	}
	v1beta1CompositionRevisionSpec.Resources = v1beta1ComposedTemplateList
	var pString *string
	if source.WriteConnectionSecretsToNamespace != nil {
		xstring := *source.WriteConnectionSecretsToNamespace
		pString = &xstring
	}
	v1beta1CompositionRevisionSpec.WriteConnectionSecretsToNamespace = pString
	var pV1beta1StoreConfigReference *v1beta1.StoreConfigReference
	if source.PublishConnectionDetailsWithStoreConfigRef != nil {
		v1beta1StoreConfigReference := c.v1StoreConfigReferenceToV1beta1StoreConfigReference(*source.PublishConnectionDetailsWithStoreConfigRef)
		pV1beta1StoreConfigReference = &v1beta1StoreConfigReference
	}
	v1beta1CompositionRevisionSpec.PublishConnectionDetailsWithStoreConfigRef = pV1beta1StoreConfigReference
	return v1beta1CompositionRevisionSpec
}
func (c *GeneratedRevisionSpecConverter) v1CombineToV1beta1Combine(source Combine) v1beta1.Combine {
	var v1beta1Combine v1beta1.Combine
	v1beta1CombineVariableList := make([]v1beta1.CombineVariable, len(source.Variables))
	for i := 0; i < len(source.Variables); i++ {
		v1beta1CombineVariableList[i] = c.v1CombineVariableToV1beta1CombineVariable(source.Variables[i])
	}
	v1beta1Combine.Variables = v1beta1CombineVariableList
	v1beta1Combine.Strategy = v1beta1.CombineStrategy(source.Strategy)
	var pV1beta1StringCombine *v1beta1.StringCombine
	if source.String != nil {
		v1beta1StringCombine := c.v1StringCombineToV1beta1StringCombine(*source.String)
		pV1beta1StringCombine = &v1beta1StringCombine
	}
	v1beta1Combine.String = pV1beta1StringCombine
	return v1beta1Combine
}
func (c *GeneratedRevisionSpecConverter) v1CombineVariableToV1beta1CombineVariable(source CombineVariable) v1beta1.CombineVariable {
	var v1beta1CombineVariable v1beta1.CombineVariable
	v1beta1CombineVariable.FromFieldPath = source.FromFieldPath
	return v1beta1CombineVariable
}
func (c *GeneratedRevisionSpecConverter) v1ComposedTemplateToV1beta1ComposedTemplate(source ComposedTemplate) v1beta1.ComposedTemplate {
	var v1beta1ComposedTemplate v1beta1.ComposedTemplate
	var pString *string
	if source.Name != nil {
		xstring := *source.Name
		pString = &xstring
	}
	v1beta1ComposedTemplate.Name = pString
	v1beta1ComposedTemplate.Base = ConvertRawExtension(source.Base)
	v1beta1PatchList := make([]v1beta1.Patch, len(source.Patches))
	for i := 0; i < len(source.Patches); i++ {
		v1beta1PatchList[i] = c.v1PatchToV1beta1Patch(source.Patches[i])
	}
	v1beta1ComposedTemplate.Patches = v1beta1PatchList
	v1beta1ConnectionDetailList := make([]v1beta1.ConnectionDetail, len(source.ConnectionDetails))
	for j := 0; j < len(source.ConnectionDetails); j++ {
		v1beta1ConnectionDetailList[j] = c.v1ConnectionDetailToV1beta1ConnectionDetail(source.ConnectionDetails[j])
	}
	v1beta1ComposedTemplate.ConnectionDetails = v1beta1ConnectionDetailList
	v1beta1ReadinessCheckList := make([]v1beta1.ReadinessCheck, len(source.ReadinessChecks))
	for k := 0; k < len(source.ReadinessChecks); k++ {
		v1beta1ReadinessCheckList[k] = c.v1ReadinessCheckToV1beta1ReadinessCheck(source.ReadinessChecks[k])
	}
	v1beta1ComposedTemplate.ReadinessChecks = v1beta1ReadinessCheckList
	return v1beta1ComposedTemplate
}
func (c *GeneratedRevisionSpecConverter) v1ConnectionDetailToV1beta1ConnectionDetail(source ConnectionDetail) v1beta1.ConnectionDetail {
	var v1beta1ConnectionDetail v1beta1.ConnectionDetail
	var pString *string
	if source.Name != nil {
		xstring := *source.Name
		pString = &xstring
	}
	v1beta1ConnectionDetail.Name = pString
	var pV1beta1ConnectionDetailType *v1beta1.ConnectionDetailType
	if source.Type != nil {
		v1beta1ConnectionDetailType := v1beta1.ConnectionDetailType(*source.Type)
		pV1beta1ConnectionDetailType = &v1beta1ConnectionDetailType
	}
	v1beta1ConnectionDetail.Type = pV1beta1ConnectionDetailType
	var pString2 *string
	if source.FromConnectionSecretKey != nil {
		xstring2 := *source.FromConnectionSecretKey
		pString2 = &xstring2
	}
	v1beta1ConnectionDetail.FromConnectionSecretKey = pString2
	var pString3 *string
	if source.FromFieldPath != nil {
		xstring3 := *source.FromFieldPath
		pString3 = &xstring3
	}
	v1beta1ConnectionDetail.FromFieldPath = pString3
	var pString4 *string
	if source.Value != nil {
		xstring4 := *source.Value
		pString4 = &xstring4
	}
	v1beta1ConnectionDetail.Value = pString4
	return v1beta1ConnectionDetail
}
func (c *GeneratedRevisionSpecConverter) v1ConvertTransformToV1beta1ConvertTransform(source ConvertTransform) v1beta1.ConvertTransform {
	var v1beta1ConvertTransform v1beta1.ConvertTransform
	v1beta1ConvertTransform.ToType = source.ToType
	return v1beta1ConvertTransform
}
func (c *GeneratedRevisionSpecConverter) v1EnvironmentConfigurationToV1beta1EnvironmentConfiguration(source EnvironmentConfiguration) v1beta1.EnvironmentConfiguration {
	var v1beta1EnvironmentConfiguration v1beta1.EnvironmentConfiguration
	v1beta1EnvironmentSourceList := make([]v1beta1.EnvironmentSource, len(source.EnvironmentConfigs))
	for i := 0; i < len(source.EnvironmentConfigs); i++ {
		v1beta1EnvironmentSourceList[i] = c.v1EnvironmentSourceToV1beta1EnvironmentSource(source.EnvironmentConfigs[i])
	}
	v1beta1EnvironmentConfiguration.EnvironmentConfigs = v1beta1EnvironmentSourceList
	v1beta1EnvironmentPatchList := make([]v1beta1.EnvironmentPatch, len(source.Patches))
	for j := 0; j < len(source.Patches); j++ {
		v1beta1EnvironmentPatchList[j] = c.v1EnvironmentPatchToV1beta1EnvironmentPatch(source.Patches[j])
	}
	v1beta1EnvironmentConfiguration.Patches = v1beta1EnvironmentPatchList
	return v1beta1EnvironmentConfiguration
}
func (c *GeneratedRevisionSpecConverter) v1EnvironmentPatchToV1beta1EnvironmentPatch(source EnvironmentPatch) v1beta1.EnvironmentPatch {
	var v1beta1EnvironmentPatch v1beta1.EnvironmentPatch
	v1beta1EnvironmentPatch.Type = v1beta1.PatchType(source.Type)
	var pString *string
	if source.FromFieldPath != nil {
		xstring := *source.FromFieldPath
		pString = &xstring
	}
	v1beta1EnvironmentPatch.FromFieldPath = pString
	var pV1beta1Combine *v1beta1.Combine
	if source.Combine != nil {
		v1beta1Combine := c.v1CombineToV1beta1Combine(*source.Combine)
		pV1beta1Combine = &v1beta1Combine
	}
	v1beta1EnvironmentPatch.Combine = pV1beta1Combine
	var pString2 *string
	if source.ToFieldPath != nil {
		xstring2 := *source.ToFieldPath
		pString2 = &xstring2
	}
	v1beta1EnvironmentPatch.ToFieldPath = pString2
	v1beta1TransformList := make([]v1beta1.Transform, len(source.Transforms))
	for i := 0; i < len(source.Transforms); i++ {
		v1beta1TransformList[i] = c.v1TransformToV1beta1Transform(source.Transforms[i])
	}
	v1beta1EnvironmentPatch.Transforms = v1beta1TransformList
	var pV1beta1PatchPolicy *v1beta1.PatchPolicy
	if source.Policy != nil {
		v1beta1PatchPolicy := c.v1PatchPolicyToV1beta1PatchPolicy(*source.Policy)
		pV1beta1PatchPolicy = &v1beta1PatchPolicy
	}
	v1beta1EnvironmentPatch.Policy = pV1beta1PatchPolicy
	return v1beta1EnvironmentPatch
}
func (c *GeneratedRevisionSpecConverter) v1EnvironmentSourceReferenceToV1beta1EnvironmentSourceReference(source EnvironmentSourceReference) v1beta1.EnvironmentSourceReference {
	var v1beta1EnvironmentSourceReference v1beta1.EnvironmentSourceReference
	v1beta1EnvironmentSourceReference.Name = source.Name
	return v1beta1EnvironmentSourceReference
}
func (c *GeneratedRevisionSpecConverter) v1EnvironmentSourceSelectorLabelMatcherToV1beta1EnvironmentSourceSelectorLabelMatcher(source EnvironmentSourceSelectorLabelMatcher) v1beta1.EnvironmentSourceSelectorLabelMatcher {
	var v1beta1EnvironmentSourceSelectorLabelMatcher v1beta1.EnvironmentSourceSelectorLabelMatcher
	v1beta1EnvironmentSourceSelectorLabelMatcher.Type = v1beta1.EnvironmentSourceSelectorLabelMatcherType(source.Type)
	v1beta1EnvironmentSourceSelectorLabelMatcher.Key = source.Key
	var pString *string
	if source.ValueFromFieldPath != nil {
		xstring := *source.ValueFromFieldPath
		pString = &xstring
	}
	v1beta1EnvironmentSourceSelectorLabelMatcher.ValueFromFieldPath = pString
	var pString2 *string
	if source.Value != nil {
		xstring2 := *source.Value
		pString2 = &xstring2
	}
	v1beta1EnvironmentSourceSelectorLabelMatcher.Value = pString2
	return v1beta1EnvironmentSourceSelectorLabelMatcher
}
func (c *GeneratedRevisionSpecConverter) v1EnvironmentSourceSelectorToV1beta1EnvironmentSourceSelector(source EnvironmentSourceSelector) v1beta1.EnvironmentSourceSelector {
	var v1beta1EnvironmentSourceSelector v1beta1.EnvironmentSourceSelector
	v1beta1EnvironmentSourceSelectorLabelMatcherList := make([]v1beta1.EnvironmentSourceSelectorLabelMatcher, len(source.MatchLabels))
	for i := 0; i < len(source.MatchLabels); i++ {
		v1beta1EnvironmentSourceSelectorLabelMatcherList[i] = c.v1EnvironmentSourceSelectorLabelMatcherToV1beta1EnvironmentSourceSelectorLabelMatcher(source.MatchLabels[i])
	}
	v1beta1EnvironmentSourceSelector.MatchLabels = v1beta1EnvironmentSourceSelectorLabelMatcherList
	return v1beta1EnvironmentSourceSelector
}
func (c *GeneratedRevisionSpecConverter) v1EnvironmentSourceToV1beta1EnvironmentSource(source EnvironmentSource) v1beta1.EnvironmentSource {
	var v1beta1EnvironmentSource v1beta1.EnvironmentSource
	v1beta1EnvironmentSource.Type = v1beta1.EnvironmentSourceType(source.Type)
	var pV1beta1EnvironmentSourceReference *v1beta1.EnvironmentSourceReference
	if source.Ref != nil {
		v1beta1EnvironmentSourceReference := c.v1EnvironmentSourceReferenceToV1beta1EnvironmentSourceReference(*source.Ref)
		pV1beta1EnvironmentSourceReference = &v1beta1EnvironmentSourceReference
	}
	v1beta1EnvironmentSource.Ref = pV1beta1EnvironmentSourceReference
	var pV1beta1EnvironmentSourceSelector *v1beta1.EnvironmentSourceSelector
	if source.Selector != nil {
		v1beta1EnvironmentSourceSelector := c.v1EnvironmentSourceSelectorToV1beta1EnvironmentSourceSelector(*source.Selector)
		pV1beta1EnvironmentSourceSelector = &v1beta1EnvironmentSourceSelector
	}
	v1beta1EnvironmentSource.Selector = pV1beta1EnvironmentSourceSelector
	return v1beta1EnvironmentSource
}
func (c *GeneratedRevisionSpecConverter) v1JSONToV1JSON(source v1.JSON) v1.JSON {
	var v1JSON v1.JSON
	byteList := make([]uint8, len(source.Raw))
	for i := 0; i < len(source.Raw); i++ {
		byteList[i] = source.Raw[i]
	}
	v1JSON.Raw = byteList
	return v1JSON
}
func (c *GeneratedRevisionSpecConverter) v1MapTransformToV1beta1MapTransform(source MapTransform) v1beta1.MapTransform {
	var v1beta1MapTransform v1beta1.MapTransform
	mapStringV1JSON := make(map[string]v1.JSON, len(source.Pairs))
	for key, value := range source.Pairs {
		mapStringV1JSON[key] = c.v1JSONToV1JSON(value)
	}
	v1beta1MapTransform.Pairs = mapStringV1JSON
	return v1beta1MapTransform
}
func (c *GeneratedRevisionSpecConverter) v1MatchTransformPatternToV1beta1MatchTransformPattern(source MatchTransformPattern) v1beta1.MatchTransformPattern {
	var v1beta1MatchTransformPattern v1beta1.MatchTransformPattern
	v1beta1MatchTransformPattern.Type = v1beta1.MatchTransformPatternType(source.Type)
	var pString *string
	if source.Literal != nil {
		xstring := *source.Literal
		pString = &xstring
	}
	v1beta1MatchTransformPattern.Literal = pString
	var pString2 *string
	if source.Regexp != nil {
		xstring2 := *source.Regexp
		pString2 = &xstring2
	}
	v1beta1MatchTransformPattern.Regexp = pString2
	v1beta1MatchTransformPattern.Result = c.v1JSONToV1JSON(source.Result)
	return v1beta1MatchTransformPattern
}
func (c *GeneratedRevisionSpecConverter) v1MatchTransformToV1beta1MatchTransform(source MatchTransform) v1beta1.MatchTransform {
	var v1beta1MatchTransform v1beta1.MatchTransform
	v1beta1MatchTransformPatternList := make([]v1beta1.MatchTransformPattern, len(source.Patterns))
	for i := 0; i < len(source.Patterns); i++ {
		v1beta1MatchTransformPatternList[i] = c.v1MatchTransformPatternToV1beta1MatchTransformPattern(source.Patterns[i])
	}
	v1beta1MatchTransform.Patterns = v1beta1MatchTransformPatternList
	v1beta1MatchTransform.FallbackValue = c.v1JSONToV1JSON(source.FallbackValue)
	return v1beta1MatchTransform
}
func (c *GeneratedRevisionSpecConverter) v1MathTransformToV1beta1MathTransform(source MathTransform) v1beta1.MathTransform {
	var v1beta1MathTransform v1beta1.MathTransform
	var pInt64 *int64
	if source.Multiply != nil {
		xint64 := *source.Multiply
		pInt64 = &xint64
	}
	v1beta1MathTransform.Multiply = pInt64
	return v1beta1MathTransform
}
func (c *GeneratedRevisionSpecConverter) v1MergeOptionsToV1MergeOptions(source v11.MergeOptions) v11.MergeOptions {
	var v1MergeOptions v11.MergeOptions
	var pBool *bool
	if source.KeepMapValues != nil {
		xbool := *source.KeepMapValues
		pBool = &xbool
	}
	v1MergeOptions.KeepMapValues = pBool
	var pBool2 *bool
	if source.AppendSlice != nil {
		xbool2 := *source.AppendSlice
		pBool2 = &xbool2
	}
	v1MergeOptions.AppendSlice = pBool2
	return v1MergeOptions
}
func (c *GeneratedRevisionSpecConverter) v1PatchPolicyToV1beta1PatchPolicy(source PatchPolicy) v1beta1.PatchPolicy {
	var v1beta1PatchPolicy v1beta1.PatchPolicy
	var pV1beta1FromFieldPathPolicy *v1beta1.FromFieldPathPolicy
	if source.FromFieldPath != nil {
		v1beta1FromFieldPathPolicy := v1beta1.FromFieldPathPolicy(*source.FromFieldPath)
		pV1beta1FromFieldPathPolicy = &v1beta1FromFieldPathPolicy
	}
	v1beta1PatchPolicy.FromFieldPath = pV1beta1FromFieldPathPolicy
	var pV1MergeOptions *v11.MergeOptions
	if source.MergeOptions != nil {
		v1MergeOptions := c.v1MergeOptionsToV1MergeOptions(*source.MergeOptions)
		pV1MergeOptions = &v1MergeOptions
	}
	v1beta1PatchPolicy.MergeOptions = pV1MergeOptions
	return v1beta1PatchPolicy
}
func (c *GeneratedRevisionSpecConverter) v1PatchSetToV1beta1PatchSet(source PatchSet) v1beta1.PatchSet {
	var v1beta1PatchSet v1beta1.PatchSet
	v1beta1PatchSet.Name = source.Name
	v1beta1PatchList := make([]v1beta1.Patch, len(source.Patches))
	for i := 0; i < len(source.Patches); i++ {
		v1beta1PatchList[i] = c.v1PatchToV1beta1Patch(source.Patches[i])
	}
	v1beta1PatchSet.Patches = v1beta1PatchList
	return v1beta1PatchSet
}
func (c *GeneratedRevisionSpecConverter) v1PatchToV1beta1Patch(source Patch) v1beta1.Patch {
	var v1beta1Patch v1beta1.Patch
	v1beta1Patch.Type = v1beta1.PatchType(source.Type)
	var pString *string
	if source.FromFieldPath != nil {
		xstring := *source.FromFieldPath
		pString = &xstring
	}
	v1beta1Patch.FromFieldPath = pString
	var pV1beta1Combine *v1beta1.Combine
	if source.Combine != nil {
		v1beta1Combine := c.v1CombineToV1beta1Combine(*source.Combine)
		pV1beta1Combine = &v1beta1Combine
	}
	v1beta1Patch.Combine = pV1beta1Combine
	var pString2 *string
	if source.ToFieldPath != nil {
		xstring2 := *source.ToFieldPath
		pString2 = &xstring2
	}
	v1beta1Patch.ToFieldPath = pString2
	var pString3 *string
	if source.PatchSetName != nil {
		xstring3 := *source.PatchSetName
		pString3 = &xstring3
	}
	v1beta1Patch.PatchSetName = pString3
	v1beta1TransformList := make([]v1beta1.Transform, len(source.Transforms))
	for i := 0; i < len(source.Transforms); i++ {
		v1beta1TransformList[i] = c.v1TransformToV1beta1Transform(source.Transforms[i])
	}
	v1beta1Patch.Transforms = v1beta1TransformList
	var pV1beta1PatchPolicy *v1beta1.PatchPolicy
	if source.Policy != nil {
		v1beta1PatchPolicy := c.v1PatchPolicyToV1beta1PatchPolicy(*source.Policy)
		pV1beta1PatchPolicy = &v1beta1PatchPolicy
	}
	v1beta1Patch.Policy = pV1beta1PatchPolicy
	return v1beta1Patch
}
func (c *GeneratedRevisionSpecConverter) v1ReadinessCheckToV1beta1ReadinessCheck(source ReadinessCheck) v1beta1.ReadinessCheck {
	var v1beta1ReadinessCheck v1beta1.ReadinessCheck
	v1beta1ReadinessCheck.Type = v1beta1.ReadinessCheckType(source.Type)
	v1beta1ReadinessCheck.FieldPath = source.FieldPath
	v1beta1ReadinessCheck.MatchString = source.MatchString
	v1beta1ReadinessCheck.MatchInteger = source.MatchInteger
	return v1beta1ReadinessCheck
}
func (c *GeneratedRevisionSpecConverter) v1StoreConfigReferenceToV1beta1StoreConfigReference(source StoreConfigReference) v1beta1.StoreConfigReference {
	var v1beta1StoreConfigReference v1beta1.StoreConfigReference
	v1beta1StoreConfigReference.Name = source.Name
	return v1beta1StoreConfigReference
}
func (c *GeneratedRevisionSpecConverter) v1StringCombineToV1beta1StringCombine(source StringCombine) v1beta1.StringCombine {
	var v1beta1StringCombine v1beta1.StringCombine
	v1beta1StringCombine.Format = source.Format
	return v1beta1StringCombine
}
func (c *GeneratedRevisionSpecConverter) v1StringTransformRegexpToV1beta1StringTransformRegexp(source StringTransformRegexp) v1beta1.StringTransformRegexp {
	var v1beta1StringTransformRegexp v1beta1.StringTransformRegexp
	v1beta1StringTransformRegexp.Match = source.Match
	var pInt *int
	if source.Group != nil {
		xint := *source.Group
		pInt = &xint
	}
	v1beta1StringTransformRegexp.Group = pInt
	return v1beta1StringTransformRegexp
}
func (c *GeneratedRevisionSpecConverter) v1StringTransformToV1beta1StringTransform(source StringTransform) v1beta1.StringTransform {
	var v1beta1StringTransform v1beta1.StringTransform
	v1beta1StringTransform.Type = v1beta1.StringTransformType(source.Type)
	var pString *string
	if source.Format != nil {
		xstring := *source.Format
		pString = &xstring
	}
	v1beta1StringTransform.Format = pString
	var pV1beta1StringConversionType *v1beta1.StringConversionType
	if source.Convert != nil {
		v1beta1StringConversionType := v1beta1.StringConversionType(*source.Convert)
		pV1beta1StringConversionType = &v1beta1StringConversionType
	}
	v1beta1StringTransform.Convert = pV1beta1StringConversionType
	var pString2 *string
	if source.Trim != nil {
		xstring2 := *source.Trim
		pString2 = &xstring2
	}
	v1beta1StringTransform.Trim = pString2
	var pV1beta1StringTransformRegexp *v1beta1.StringTransformRegexp
	if source.Regexp != nil {
		v1beta1StringTransformRegexp := c.v1StringTransformRegexpToV1beta1StringTransformRegexp(*source.Regexp)
		pV1beta1StringTransformRegexp = &v1beta1StringTransformRegexp
	}
	v1beta1StringTransform.Regexp = pV1beta1StringTransformRegexp
	return v1beta1StringTransform
}
func (c *GeneratedRevisionSpecConverter) v1TransformToV1beta1Transform(source Transform) v1beta1.Transform {
	var v1beta1Transform v1beta1.Transform
	v1beta1Transform.Type = v1beta1.TransformType(source.Type)
	var pV1beta1MathTransform *v1beta1.MathTransform
	if source.Math != nil {
		v1beta1MathTransform := c.v1MathTransformToV1beta1MathTransform(*source.Math)
		pV1beta1MathTransform = &v1beta1MathTransform
	}
	v1beta1Transform.Math = pV1beta1MathTransform
	var pV1beta1MapTransform *v1beta1.MapTransform
	if source.Map != nil {
		v1beta1MapTransform := c.v1MapTransformToV1beta1MapTransform(*source.Map)
		pV1beta1MapTransform = &v1beta1MapTransform
	}
	v1beta1Transform.Map = pV1beta1MapTransform
	var pV1beta1MatchTransform *v1beta1.MatchTransform
	if source.Match != nil {
		v1beta1MatchTransform := c.v1MatchTransformToV1beta1MatchTransform(*source.Match)
		pV1beta1MatchTransform = &v1beta1MatchTransform
	}
	v1beta1Transform.Match = pV1beta1MatchTransform
	var pV1beta1StringTransform *v1beta1.StringTransform
	if source.String != nil {
		v1beta1StringTransform := c.v1StringTransformToV1beta1StringTransform(*source.String)
		pV1beta1StringTransform = &v1beta1StringTransform
	}
	v1beta1Transform.String = pV1beta1StringTransform
	var pV1beta1ConvertTransform *v1beta1.ConvertTransform
	if source.Convert != nil {
		v1beta1ConvertTransform := c.v1ConvertTransformToV1beta1ConvertTransform(*source.Convert)
		pV1beta1ConvertTransform = &v1beta1ConvertTransform
	}
	v1beta1Transform.Convert = pV1beta1ConvertTransform
	return v1beta1Transform
}
func (c *GeneratedRevisionSpecConverter) v1TypeReferenceToV1beta1TypeReference(source TypeReference) v1beta1.TypeReference {
	var v1beta1TypeReference v1beta1.TypeReference
	v1beta1TypeReference.APIVersion = source.APIVersion
	v1beta1TypeReference.Kind = source.Kind
	return v1beta1TypeReference
}
func (c *GeneratedRevisionSpecConverter) v1beta1CombineToV1Combine(source v1beta1.Combine) Combine {
	var v1Combine Combine
	v1CombineVariableList := make([]CombineVariable, len(source.Variables))
	for i := 0; i < len(source.Variables); i++ {
		v1CombineVariableList[i] = c.v1beta1CombineVariableToV1CombineVariable(source.Variables[i])
	}
	v1Combine.Variables = v1CombineVariableList
	v1Combine.Strategy = CombineStrategy(source.Strategy)
	var pV1StringCombine *StringCombine
	if source.String != nil {
		v1StringCombine := c.v1beta1StringCombineToV1StringCombine(*source.String)
		pV1StringCombine = &v1StringCombine
	}
	v1Combine.String = pV1StringCombine
	return v1Combine
}
func (c *GeneratedRevisionSpecConverter) v1beta1CombineVariableToV1CombineVariable(source v1beta1.CombineVariable) CombineVariable {
	var v1CombineVariable CombineVariable
	v1CombineVariable.FromFieldPath = source.FromFieldPath
	return v1CombineVariable
}
func (c *GeneratedRevisionSpecConverter) v1beta1ComposedTemplateToV1ComposedTemplate(source v1beta1.ComposedTemplate) ComposedTemplate {
	var v1ComposedTemplate ComposedTemplate
	var pString *string
	if source.Name != nil {
		xstring := *source.Name
		pString = &xstring
	}
	v1ComposedTemplate.Name = pString
	v1ComposedTemplate.Base = ConvertRawExtension(source.Base)
	v1PatchList := make([]Patch, len(source.Patches))
	for i := 0; i < len(source.Patches); i++ {
		v1PatchList[i] = c.v1beta1PatchToV1Patch(source.Patches[i])
	}
	v1ComposedTemplate.Patches = v1PatchList
	v1ConnectionDetailList := make([]ConnectionDetail, len(source.ConnectionDetails))
	for j := 0; j < len(source.ConnectionDetails); j++ {
		v1ConnectionDetailList[j] = c.v1beta1ConnectionDetailToV1ConnectionDetail(source.ConnectionDetails[j])
	}
	v1ComposedTemplate.ConnectionDetails = v1ConnectionDetailList
	v1ReadinessCheckList := make([]ReadinessCheck, len(source.ReadinessChecks))
	for k := 0; k < len(source.ReadinessChecks); k++ {
		v1ReadinessCheckList[k] = c.v1beta1ReadinessCheckToV1ReadinessCheck(source.ReadinessChecks[k])
	}
	v1ComposedTemplate.ReadinessChecks = v1ReadinessCheckList
	return v1ComposedTemplate
}
func (c *GeneratedRevisionSpecConverter) v1beta1ConnectionDetailToV1ConnectionDetail(source v1beta1.ConnectionDetail) ConnectionDetail {
	var v1ConnectionDetail ConnectionDetail
	var pString *string
	if source.Name != nil {
		xstring := *source.Name
		pString = &xstring
	}
	v1ConnectionDetail.Name = pString
	var pV1ConnectionDetailType *ConnectionDetailType
	if source.Type != nil {
		v1ConnectionDetailType := ConnectionDetailType(*source.Type)
		pV1ConnectionDetailType = &v1ConnectionDetailType
	}
	v1ConnectionDetail.Type = pV1ConnectionDetailType
	var pString2 *string
	if source.FromConnectionSecretKey != nil {
		xstring2 := *source.FromConnectionSecretKey
		pString2 = &xstring2
	}
	v1ConnectionDetail.FromConnectionSecretKey = pString2
	var pString3 *string
	if source.FromFieldPath != nil {
		xstring3 := *source.FromFieldPath
		pString3 = &xstring3
	}
	v1ConnectionDetail.FromFieldPath = pString3
	var pString4 *string
	if source.Value != nil {
		xstring4 := *source.Value
		pString4 = &xstring4
	}
	v1ConnectionDetail.Value = pString4
	return v1ConnectionDetail
}
func (c *GeneratedRevisionSpecConverter) v1beta1ConvertTransformToV1ConvertTransform(source v1beta1.ConvertTransform) ConvertTransform {
	var v1ConvertTransform ConvertTransform
	v1ConvertTransform.ToType = source.ToType
	return v1ConvertTransform
}
func (c *GeneratedRevisionSpecConverter) v1beta1EnvironmentConfigurationToV1EnvironmentConfiguration(source v1beta1.EnvironmentConfiguration) EnvironmentConfiguration {
	var v1EnvironmentConfiguration EnvironmentConfiguration
	v1EnvironmentSourceList := make([]EnvironmentSource, len(source.EnvironmentConfigs))
	for i := 0; i < len(source.EnvironmentConfigs); i++ {
		v1EnvironmentSourceList[i] = c.v1beta1EnvironmentSourceToV1EnvironmentSource(source.EnvironmentConfigs[i])
	}
	v1EnvironmentConfiguration.EnvironmentConfigs = v1EnvironmentSourceList
	v1EnvironmentPatchList := make([]EnvironmentPatch, len(source.Patches))
	for j := 0; j < len(source.Patches); j++ {
		v1EnvironmentPatchList[j] = c.v1beta1EnvironmentPatchToV1EnvironmentPatch(source.Patches[j])
	}
	v1EnvironmentConfiguration.Patches = v1EnvironmentPatchList
	return v1EnvironmentConfiguration
}
func (c *GeneratedRevisionSpecConverter) v1beta1EnvironmentPatchToV1EnvironmentPatch(source v1beta1.EnvironmentPatch) EnvironmentPatch {
	var v1EnvironmentPatch EnvironmentPatch
	v1EnvironmentPatch.Type = PatchType(source.Type)
	var pString *string
	if source.FromFieldPath != nil {
		xstring := *source.FromFieldPath
		pString = &xstring
	}
	v1EnvironmentPatch.FromFieldPath = pString
	var pV1Combine *Combine
	if source.Combine != nil {
		v1Combine := c.v1beta1CombineToV1Combine(*source.Combine)
		pV1Combine = &v1Combine
	}
	v1EnvironmentPatch.Combine = pV1Combine
	var pString2 *string
	if source.ToFieldPath != nil {
		xstring2 := *source.ToFieldPath
		pString2 = &xstring2
	}
	v1EnvironmentPatch.ToFieldPath = pString2
	v1TransformList := make([]Transform, len(source.Transforms))
	for i := 0; i < len(source.Transforms); i++ {
		v1TransformList[i] = c.v1beta1TransformToV1Transform(source.Transforms[i])
	}
	v1EnvironmentPatch.Transforms = v1TransformList
	var pV1PatchPolicy *PatchPolicy
	if source.Policy != nil {
		v1PatchPolicy := c.v1beta1PatchPolicyToV1PatchPolicy(*source.Policy)
		pV1PatchPolicy = &v1PatchPolicy
	}
	v1EnvironmentPatch.Policy = pV1PatchPolicy
	return v1EnvironmentPatch
}
func (c *GeneratedRevisionSpecConverter) v1beta1EnvironmentSourceReferenceToV1EnvironmentSourceReference(source v1beta1.EnvironmentSourceReference) EnvironmentSourceReference {
	var v1EnvironmentSourceReference EnvironmentSourceReference
	v1EnvironmentSourceReference.Name = source.Name
	return v1EnvironmentSourceReference
}
func (c *GeneratedRevisionSpecConverter) v1beta1EnvironmentSourceSelectorLabelMatcherToV1EnvironmentSourceSelectorLabelMatcher(source v1beta1.EnvironmentSourceSelectorLabelMatcher) EnvironmentSourceSelectorLabelMatcher {
	var v1EnvironmentSourceSelectorLabelMatcher EnvironmentSourceSelectorLabelMatcher
	v1EnvironmentSourceSelectorLabelMatcher.Type = EnvironmentSourceSelectorLabelMatcherType(source.Type)
	v1EnvironmentSourceSelectorLabelMatcher.Key = source.Key
	var pString *string
	if source.ValueFromFieldPath != nil {
		xstring := *source.ValueFromFieldPath
		pString = &xstring
	}
	v1EnvironmentSourceSelectorLabelMatcher.ValueFromFieldPath = pString
	var pString2 *string
	if source.Value != nil {
		xstring2 := *source.Value
		pString2 = &xstring2
	}
	v1EnvironmentSourceSelectorLabelMatcher.Value = pString2
	return v1EnvironmentSourceSelectorLabelMatcher
}
func (c *GeneratedRevisionSpecConverter) v1beta1EnvironmentSourceSelectorToV1EnvironmentSourceSelector(source v1beta1.EnvironmentSourceSelector) EnvironmentSourceSelector {
	var v1EnvironmentSourceSelector EnvironmentSourceSelector
	v1EnvironmentSourceSelectorLabelMatcherList := make([]EnvironmentSourceSelectorLabelMatcher, len(source.MatchLabels))
	for i := 0; i < len(source.MatchLabels); i++ {
		v1EnvironmentSourceSelectorLabelMatcherList[i] = c.v1beta1EnvironmentSourceSelectorLabelMatcherToV1EnvironmentSourceSelectorLabelMatcher(source.MatchLabels[i])
	}
	v1EnvironmentSourceSelector.MatchLabels = v1EnvironmentSourceSelectorLabelMatcherList
	return v1EnvironmentSourceSelector
}
func (c *GeneratedRevisionSpecConverter) v1beta1EnvironmentSourceToV1EnvironmentSource(source v1beta1.EnvironmentSource) EnvironmentSource {
	var v1EnvironmentSource EnvironmentSource
	v1EnvironmentSource.Type = EnvironmentSourceType(source.Type)
	var pV1EnvironmentSourceReference *EnvironmentSourceReference
	if source.Ref != nil {
		v1EnvironmentSourceReference := c.v1beta1EnvironmentSourceReferenceToV1EnvironmentSourceReference(*source.Ref)
		pV1EnvironmentSourceReference = &v1EnvironmentSourceReference
	}
	v1EnvironmentSource.Ref = pV1EnvironmentSourceReference
	var pV1EnvironmentSourceSelector *EnvironmentSourceSelector
	if source.Selector != nil {
		v1EnvironmentSourceSelector := c.v1beta1EnvironmentSourceSelectorToV1EnvironmentSourceSelector(*source.Selector)
		pV1EnvironmentSourceSelector = &v1EnvironmentSourceSelector
	}
	v1EnvironmentSource.Selector = pV1EnvironmentSourceSelector
	return v1EnvironmentSource
}
func (c *GeneratedRevisionSpecConverter) v1beta1MapTransformToV1MapTransform(source v1beta1.MapTransform) MapTransform {
	var v1MapTransform MapTransform
	mapStringV1JSON := make(map[string]v1.JSON, len(source.Pairs))
	for key, value := range source.Pairs {
		mapStringV1JSON[key] = c.v1JSONToV1JSON(value)
	}
	v1MapTransform.Pairs = mapStringV1JSON
	return v1MapTransform
}
func (c *GeneratedRevisionSpecConverter) v1beta1MatchTransformPatternToV1MatchTransformPattern(source v1beta1.MatchTransformPattern) MatchTransformPattern {
	var v1MatchTransformPattern MatchTransformPattern
	v1MatchTransformPattern.Type = MatchTransformPatternType(source.Type)
	var pString *string
	if source.Literal != nil {
		xstring := *source.Literal
		pString = &xstring
	}
	v1MatchTransformPattern.Literal = pString
	var pString2 *string
	if source.Regexp != nil {
		xstring2 := *source.Regexp
		pString2 = &xstring2
	}
	v1MatchTransformPattern.Regexp = pString2
	v1MatchTransformPattern.Result = c.v1JSONToV1JSON(source.Result)
	return v1MatchTransformPattern
}
func (c *GeneratedRevisionSpecConverter) v1beta1MatchTransformToV1MatchTransform(source v1beta1.MatchTransform) MatchTransform {
	var v1MatchTransform MatchTransform
	v1MatchTransformPatternList := make([]MatchTransformPattern, len(source.Patterns))
	for i := 0; i < len(source.Patterns); i++ {
		v1MatchTransformPatternList[i] = c.v1beta1MatchTransformPatternToV1MatchTransformPattern(source.Patterns[i])
	}
	v1MatchTransform.Patterns = v1MatchTransformPatternList
	v1MatchTransform.FallbackValue = c.v1JSONToV1JSON(source.FallbackValue)
	return v1MatchTransform
}
func (c *GeneratedRevisionSpecConverter) v1beta1MathTransformToV1MathTransform(source v1beta1.MathTransform) MathTransform {
	var v1MathTransform MathTransform
	var pInt64 *int64
	if source.Multiply != nil {
		xint64 := *source.Multiply
		pInt64 = &xint64
	}
	v1MathTransform.Multiply = pInt64
	return v1MathTransform
}
func (c *GeneratedRevisionSpecConverter) v1beta1PatchPolicyToV1PatchPolicy(source v1beta1.PatchPolicy) PatchPolicy {
	var v1PatchPolicy PatchPolicy
	var pV1FromFieldPathPolicy *FromFieldPathPolicy
	if source.FromFieldPath != nil {
		v1FromFieldPathPolicy := FromFieldPathPolicy(*source.FromFieldPath)
		pV1FromFieldPathPolicy = &v1FromFieldPathPolicy
	}
	v1PatchPolicy.FromFieldPath = pV1FromFieldPathPolicy
	var pV1MergeOptions *v11.MergeOptions
	if source.MergeOptions != nil {
		v1MergeOptions := c.v1MergeOptionsToV1MergeOptions(*source.MergeOptions)
		pV1MergeOptions = &v1MergeOptions
	}
	v1PatchPolicy.MergeOptions = pV1MergeOptions
	return v1PatchPolicy
}
func (c *GeneratedRevisionSpecConverter) v1beta1PatchSetToV1PatchSet(source v1beta1.PatchSet) PatchSet {
	var v1PatchSet PatchSet
	v1PatchSet.Name = source.Name
	v1PatchList := make([]Patch, len(source.Patches))
	for i := 0; i < len(source.Patches); i++ {
		v1PatchList[i] = c.v1beta1PatchToV1Patch(source.Patches[i])
	}
	v1PatchSet.Patches = v1PatchList
	return v1PatchSet
}
func (c *GeneratedRevisionSpecConverter) v1beta1PatchToV1Patch(source v1beta1.Patch) Patch {
	var v1Patch Patch
	v1Patch.Type = PatchType(source.Type)
	var pString *string
	if source.FromFieldPath != nil {
		xstring := *source.FromFieldPath
		pString = &xstring
	}
	v1Patch.FromFieldPath = pString
	var pV1Combine *Combine
	if source.Combine != nil {
		v1Combine := c.v1beta1CombineToV1Combine(*source.Combine)
		pV1Combine = &v1Combine
	}
	v1Patch.Combine = pV1Combine
	var pString2 *string
	if source.ToFieldPath != nil {
		xstring2 := *source.ToFieldPath
		pString2 = &xstring2
	}
	v1Patch.ToFieldPath = pString2
	var pString3 *string
	if source.PatchSetName != nil {
		xstring3 := *source.PatchSetName
		pString3 = &xstring3
	}
	v1Patch.PatchSetName = pString3
	v1TransformList := make([]Transform, len(source.Transforms))
	for i := 0; i < len(source.Transforms); i++ {
		v1TransformList[i] = c.v1beta1TransformToV1Transform(source.Transforms[i])
	}
	v1Patch.Transforms = v1TransformList
	var pV1PatchPolicy *PatchPolicy
	if source.Policy != nil {
		v1PatchPolicy := c.v1beta1PatchPolicyToV1PatchPolicy(*source.Policy)
		pV1PatchPolicy = &v1PatchPolicy
	}
	v1Patch.Policy = pV1PatchPolicy
	return v1Patch
}
func (c *GeneratedRevisionSpecConverter) v1beta1ReadinessCheckToV1ReadinessCheck(source v1beta1.ReadinessCheck) ReadinessCheck {
	var v1ReadinessCheck ReadinessCheck
	v1ReadinessCheck.Type = ReadinessCheckType(source.Type)
	v1ReadinessCheck.FieldPath = source.FieldPath
	v1ReadinessCheck.MatchString = source.MatchString
	v1ReadinessCheck.MatchInteger = source.MatchInteger
	return v1ReadinessCheck
}
func (c *GeneratedRevisionSpecConverter) v1beta1StoreConfigReferenceToV1StoreConfigReference(source v1beta1.StoreConfigReference) StoreConfigReference {
	var v1StoreConfigReference StoreConfigReference
	v1StoreConfigReference.Name = source.Name
	return v1StoreConfigReference
}
func (c *GeneratedRevisionSpecConverter) v1beta1StringCombineToV1StringCombine(source v1beta1.StringCombine) StringCombine {
	var v1StringCombine StringCombine
	v1StringCombine.Format = source.Format
	return v1StringCombine
}
func (c *GeneratedRevisionSpecConverter) v1beta1StringTransformRegexpToV1StringTransformRegexp(source v1beta1.StringTransformRegexp) StringTransformRegexp {
	var v1StringTransformRegexp StringTransformRegexp
	v1StringTransformRegexp.Match = source.Match
	var pInt *int
	if source.Group != nil {
		xint := *source.Group
		pInt = &xint
	}
	v1StringTransformRegexp.Group = pInt
	return v1StringTransformRegexp
}
func (c *GeneratedRevisionSpecConverter) v1beta1StringTransformToV1StringTransform(source v1beta1.StringTransform) StringTransform {
	var v1StringTransform StringTransform
	v1StringTransform.Type = StringTransformType(source.Type)
	var pString *string
	if source.Format != nil {
		xstring := *source.Format
		pString = &xstring
	}
	v1StringTransform.Format = pString
	var pV1StringConversionType *StringConversionType
	if source.Convert != nil {
		v1StringConversionType := StringConversionType(*source.Convert)
		pV1StringConversionType = &v1StringConversionType
	}
	v1StringTransform.Convert = pV1StringConversionType
	var pString2 *string
	if source.Trim != nil {
		xstring2 := *source.Trim
		pString2 = &xstring2
	}
	v1StringTransform.Trim = pString2
	var pV1StringTransformRegexp *StringTransformRegexp
	if source.Regexp != nil {
		v1StringTransformRegexp := c.v1beta1StringTransformRegexpToV1StringTransformRegexp(*source.Regexp)
		pV1StringTransformRegexp = &v1StringTransformRegexp
	}
	v1StringTransform.Regexp = pV1StringTransformRegexp
	return v1StringTransform
}
func (c *GeneratedRevisionSpecConverter) v1beta1TransformToV1Transform(source v1beta1.Transform) Transform {
	var v1Transform Transform
	v1Transform.Type = TransformType(source.Type)
	var pV1MathTransform *MathTransform
	if source.Math != nil {
		v1MathTransform := c.v1beta1MathTransformToV1MathTransform(*source.Math)
		pV1MathTransform = &v1MathTransform
	}
	v1Transform.Math = pV1MathTransform
	var pV1MapTransform *MapTransform
	if source.Map != nil {
		v1MapTransform := c.v1beta1MapTransformToV1MapTransform(*source.Map)
		pV1MapTransform = &v1MapTransform
	}
	v1Transform.Map = pV1MapTransform
	var pV1MatchTransform *MatchTransform
	if source.Match != nil {
		v1MatchTransform := c.v1beta1MatchTransformToV1MatchTransform(*source.Match)
		pV1MatchTransform = &v1MatchTransform
	}
	v1Transform.Match = pV1MatchTransform
	var pV1StringTransform *StringTransform
	if source.String != nil {
		v1StringTransform := c.v1beta1StringTransformToV1StringTransform(*source.String)
		pV1StringTransform = &v1StringTransform
	}
	v1Transform.String = pV1StringTransform
	var pV1ConvertTransform *ConvertTransform
	if source.Convert != nil {
		v1ConvertTransform := c.v1beta1ConvertTransformToV1ConvertTransform(*source.Convert)
		pV1ConvertTransform = &v1ConvertTransform
	}
	v1Transform.Convert = pV1ConvertTransform
	return v1Transform
}
func (c *GeneratedRevisionSpecConverter) v1beta1TypeReferenceToV1TypeReference(source v1beta1.TypeReference) TypeReference {
	var v1TypeReference TypeReference
	v1TypeReference.APIVersion = source.APIVersion
	v1TypeReference.Kind = source.Kind
	return v1TypeReference
}
