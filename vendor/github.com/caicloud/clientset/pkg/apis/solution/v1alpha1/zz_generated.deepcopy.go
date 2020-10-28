// +build !ignore_autogenerated

/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CongestionArgs) DeepCopyInto(out *CongestionArgs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CongestionArgs.
func (in *CongestionArgs) DeepCopy() *CongestionArgs {
	if in == nil {
		return nil
	}
	out := new(CongestionArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Coordinate) DeepCopyInto(out *Coordinate) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Coordinate.
func (in *Coordinate) DeepCopy() *Coordinate {
	if in == nil {
		return nil
	}
	out := new(Coordinate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomTemplateOCR) DeepCopyInto(out *CustomTemplateOCR) {
	*out = *in
	if in.References != nil {
		in, out := &in.References, &out.References
		*out = make([]Identification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]Identification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomTemplateOCR.
func (in *CustomTemplateOCR) DeepCopy() *CustomTemplateOCR {
	if in == nil {
		return nil
	}
	out := new(CustomTemplateOCR)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrameImage) DeepCopyInto(out *FrameImage) {
	*out = *in
	out.Size = in.Size
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrameImage.
func (in *FrameImage) DeepCopy() *FrameImage {
	if in == nil {
		return nil
	}
	out := new(FrameImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identification) DeepCopyInto(out *Identification) {
	*out = *in
	if in.Coordinates != nil {
		in, out := &in.Coordinates, &out.Coordinates
		*out = make([]Coordinate, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identification.
func (in *Identification) DeepCopy() *Identification {
	if in == nil {
		return nil
	}
	out := new(Identification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSize) DeepCopyInto(out *ImageSize) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSize.
func (in *ImageSize) DeepCopy() *ImageSize {
	if in == nil {
		return nil
	}
	out := new(ImageSize)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OccupancyArgs) DeepCopyInto(out *OccupancyArgs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OccupancyArgs.
func (in *OccupancyArgs) DeepCopy() *OccupancyArgs {
	if in == nil {
		return nil
	}
	out := new(OccupancyArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Solution) DeepCopyInto(out *Solution) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Solution.
func (in *Solution) DeepCopy() *Solution {
	if in == nil {
		return nil
	}
	out := new(Solution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Solution) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolutionConfig) DeepCopyInto(out *SolutionConfig) {
	*out = *in
	if in.SurveillanceVideoAnalysis != nil {
		in, out := &in.SurveillanceVideoAnalysis, &out.SurveillanceVideoAnalysis
		*out = new(SurveillanceVideoAnalysis)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomTemplateOCR != nil {
		in, out := &in.CustomTemplateOCR, &out.CustomTemplateOCR
		*out = new(CustomTemplateOCR)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolutionConfig.
func (in *SolutionConfig) DeepCopy() *SolutionConfig {
	if in == nil {
		return nil
	}
	out := new(SolutionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolutionList) DeepCopyInto(out *SolutionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Solution, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolutionList.
func (in *SolutionList) DeepCopy() *SolutionList {
	if in == nil {
		return nil
	}
	out := new(SolutionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SolutionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolutionSpec) DeepCopyInto(out *SolutionSpec) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolutionSpec.
func (in *SolutionSpec) DeepCopy() *SolutionSpec {
	if in == nil {
		return nil
	}
	out := new(SolutionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SurveillanceVideoAnalysis) DeepCopyInto(out *SurveillanceVideoAnalysis) {
	*out = *in
	if in.ObjectClass != nil {
		in, out := &in.ObjectClass, &out.ObjectClass
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.FrameImage = in.FrameImage
	if in.ROI != nil {
		in, out := &in.ROI, &out.ROI
		*out = make([]Coordinate, len(*in))
		copy(*out, *in)
	}
	if in.FileList != nil {
		in, out := &in.FileList, &out.FileList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OccupancyArgs != nil {
		in, out := &in.OccupancyArgs, &out.OccupancyArgs
		*out = new(OccupancyArgs)
		**out = **in
	}
	if in.CongestionArgs != nil {
		in, out := &in.CongestionArgs, &out.CongestionArgs
		*out = new(CongestionArgs)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SurveillanceVideoAnalysis.
func (in *SurveillanceVideoAnalysis) DeepCopy() *SurveillanceVideoAnalysis {
	if in == nil {
		return nil
	}
	out := new(SurveillanceVideoAnalysis)
	in.DeepCopyInto(out)
	return out
}
