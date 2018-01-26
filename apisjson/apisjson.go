// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// THIS FILE IS AUTOMATICALLY GENERATED.

package apisjson_v1

import (
	"fmt"
	"github.com/googleapis/gnostic/compiler"
	"gopkg.in/yaml.v2"
	"regexp"
	"strings"
)

// Version returns the package name (and OpenAPI version).
func Version() string {
	return "apisjson_v1"
}

// NewAny creates an object of type Any if possible, returning an error if not.
func NewAny(in interface{}, context *compiler.Context) (*Any, error) {
	errors := make([]error, 0)
	x := &Any{}
	bytes, _ := yaml.Marshal(in)
	x.Yaml = string(bytes)
	return x, compiler.NewErrorGroupOrNil(errors)
}

// NewApi creates an object of type Api if possible, returning an error if not.
func NewApi(in interface{}, context *compiler.Context) (*Api, error) {
	errors := make([]error, 0)
	x := &Api{}
	m, ok := compiler.UnpackMap(in)
	if !ok {
		message := fmt.Sprintf("has unexpected value: %+v (%T)", in, in)
		errors = append(errors, compiler.NewError(context, message))
	} else {
		requiredKeys := []string{"description", "name"}
		missingKeys := compiler.MissingKeysInMap(m, requiredKeys)
		if len(missingKeys) > 0 {
			message := fmt.Sprintf("is missing required %s: %+v", compiler.PluralProperties(len(missingKeys)), strings.Join(missingKeys, ", "))
			errors = append(errors, compiler.NewError(context, message))
		}
		allowedKeys := []string{"baseURL", "contact", "description", "humanURL", "image", "name", "properties", "tags", "version"}
		var allowedPatterns []*regexp.Regexp
		invalidKeys := compiler.InvalidKeysInMap(m, allowedKeys, allowedPatterns)
		if len(invalidKeys) > 0 {
			message := fmt.Sprintf("has invalid %s: %+v", compiler.PluralProperties(len(invalidKeys)), strings.Join(invalidKeys, ", "))
			errors = append(errors, compiler.NewError(context, message))
		}
		// string name = 1;
		v1 := compiler.MapValueForKey(m, "name")
		if v1 != nil {
			x.Name, ok = v1.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for name: %+v (%T)", v1, v1)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string description = 2;
		v2 := compiler.MapValueForKey(m, "description")
		if v2 != nil {
			x.Description, ok = v2.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for description: %+v (%T)", v2, v2)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string image = 3;
		v3 := compiler.MapValueForKey(m, "image")
		if v3 != nil {
			x.Image, ok = v3.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for image: %+v (%T)", v3, v3)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string human_u_r_l = 4;
		v4 := compiler.MapValueForKey(m, "humanURL")
		if v4 != nil {
			x.HumanURL, ok = v4.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for humanURL: %+v (%T)", v4, v4)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string base_u_r_l = 5;
		v5 := compiler.MapValueForKey(m, "baseURL")
		if v5 != nil {
			x.BaseURL, ok = v5.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for baseURL: %+v (%T)", v5, v5)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string version = 6;
		v6 := compiler.MapValueForKey(m, "version")
		if v6 != nil {
			x.Version, ok = v6.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for version: %+v (%T)", v6, v6)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// repeated string tags = 7;
		v7 := compiler.MapValueForKey(m, "tags")
		if v7 != nil {
			v, ok := v7.([]interface{})
			if ok {
				x.Tags = compiler.ConvertInterfaceArrayToStringArray(v)
			} else {
				message := fmt.Sprintf("has unexpected value for tags: %+v (%T)", v7, v7)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// repeated Property properties = 8;
		v8 := compiler.MapValueForKey(m, "properties")
		if v8 != nil {
			// repeated Property
			x.Properties = make([]*Property, 0)
			a, ok := v8.([]interface{})
			if ok {
				for _, item := range a {
					y, err := NewProperty(item, compiler.NewContext("properties", context))
					if err != nil {
						errors = append(errors, err)
					}
					x.Properties = append(x.Properties, y)
				}
			}
		}
		// repeated Maintainer contact = 9;
		v9 := compiler.MapValueForKey(m, "contact")
		if v9 != nil {
			// repeated Maintainer
			x.Contact = make([]*Maintainer, 0)
			a, ok := v9.([]interface{})
			if ok {
				for _, item := range a {
					y, err := NewMaintainer(item, compiler.NewContext("contact", context))
					if err != nil {
						errors = append(errors, err)
					}
					x.Contact = append(x.Contact, y)
				}
			}
		}
	}
	return x, compiler.NewErrorGroupOrNil(errors)
}

// NewDocument creates an object of type Document if possible, returning an error if not.
func NewDocument(in interface{}, context *compiler.Context) (*Document, error) {
	errors := make([]error, 0)
	x := &Document{}
	m, ok := compiler.UnpackMap(in)
	if !ok {
		message := fmt.Sprintf("has unexpected value: %+v (%T)", in, in)
		errors = append(errors, compiler.NewError(context, message))
	} else {
		requiredKeys := []string{"created", "description", "modified", "name", "specificationVersion", "url"}
		missingKeys := compiler.MissingKeysInMap(m, requiredKeys)
		if len(missingKeys) > 0 {
			message := fmt.Sprintf("is missing required %s: %+v", compiler.PluralProperties(len(missingKeys)), strings.Join(missingKeys, ", "))
			errors = append(errors, compiler.NewError(context, message))
		}
		allowedKeys := []string{"apis", "created", "description", "image", "include", "maintainers", "modified", "name", "specificationVersion", "tags", "url"}
		var allowedPatterns []*regexp.Regexp
		invalidKeys := compiler.InvalidKeysInMap(m, allowedKeys, allowedPatterns)
		if len(invalidKeys) > 0 {
			message := fmt.Sprintf("has invalid %s: %+v", compiler.PluralProperties(len(invalidKeys)), strings.Join(invalidKeys, ", "))
			errors = append(errors, compiler.NewError(context, message))
		}
		// string name = 1;
		v1 := compiler.MapValueForKey(m, "name")
		if v1 != nil {
			x.Name, ok = v1.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for name: %+v (%T)", v1, v1)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string description = 2;
		v2 := compiler.MapValueForKey(m, "description")
		if v2 != nil {
			x.Description, ok = v2.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for description: %+v (%T)", v2, v2)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string image = 3;
		v3 := compiler.MapValueForKey(m, "image")
		if v3 != nil {
			x.Image, ok = v3.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for image: %+v (%T)", v3, v3)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string url = 4;
		v4 := compiler.MapValueForKey(m, "url")
		if v4 != nil {
			x.Url, ok = v4.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for url: %+v (%T)", v4, v4)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// repeated string tags = 5;
		v5 := compiler.MapValueForKey(m, "tags")
		if v5 != nil {
			v, ok := v5.([]interface{})
			if ok {
				x.Tags = compiler.ConvertInterfaceArrayToStringArray(v)
			} else {
				message := fmt.Sprintf("has unexpected value for tags: %+v (%T)", v5, v5)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string created = 6;
		v6 := compiler.MapValueForKey(m, "created")
		if v6 != nil {
			x.Created, ok = v6.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for created: %+v (%T)", v6, v6)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string modified = 7;
		v7 := compiler.MapValueForKey(m, "modified")
		if v7 != nil {
			x.Modified, ok = v7.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for modified: %+v (%T)", v7, v7)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string specification_version = 8;
		v8 := compiler.MapValueForKey(m, "specificationVersion")
		if v8 != nil {
			x.SpecificationVersion, ok = v8.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for specificationVersion: %+v (%T)", v8, v8)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// repeated Api apis = 9;
		v9 := compiler.MapValueForKey(m, "apis")
		if v9 != nil {
			// repeated Api
			x.Apis = make([]*Api, 0)
			a, ok := v9.([]interface{})
			if ok {
				for _, item := range a {
					y, err := NewApi(item, compiler.NewContext("apis", context))
					if err != nil {
						errors = append(errors, err)
					}
					x.Apis = append(x.Apis, y)
				}
			}
		}
		// repeated Include include = 10;
		v10 := compiler.MapValueForKey(m, "include")
		if v10 != nil {
			// repeated Include
			x.Include = make([]*Include, 0)
			a, ok := v10.([]interface{})
			if ok {
				for _, item := range a {
					y, err := NewInclude(item, compiler.NewContext("include", context))
					if err != nil {
						errors = append(errors, err)
					}
					x.Include = append(x.Include, y)
				}
			}
		}
		// repeated Maintainer maintainers = 11;
		v11 := compiler.MapValueForKey(m, "maintainers")
		if v11 != nil {
			// repeated Maintainer
			x.Maintainers = make([]*Maintainer, 0)
			a, ok := v11.([]interface{})
			if ok {
				for _, item := range a {
					y, err := NewMaintainer(item, compiler.NewContext("maintainers", context))
					if err != nil {
						errors = append(errors, err)
					}
					x.Maintainers = append(x.Maintainers, y)
				}
			}
		}
	}
	return x, compiler.NewErrorGroupOrNil(errors)
}

// NewInclude creates an object of type Include if possible, returning an error if not.
func NewInclude(in interface{}, context *compiler.Context) (*Include, error) {
	errors := make([]error, 0)
	x := &Include{}
	m, ok := compiler.UnpackMap(in)
	if !ok {
		message := fmt.Sprintf("has unexpected value: %+v (%T)", in, in)
		errors = append(errors, compiler.NewError(context, message))
	} else {
		requiredKeys := []string{"name", "url"}
		missingKeys := compiler.MissingKeysInMap(m, requiredKeys)
		if len(missingKeys) > 0 {
			message := fmt.Sprintf("is missing required %s: %+v", compiler.PluralProperties(len(missingKeys)), strings.Join(missingKeys, ", "))
			errors = append(errors, compiler.NewError(context, message))
		}
		allowedKeys := []string{"name", "url"}
		var allowedPatterns []*regexp.Regexp
		invalidKeys := compiler.InvalidKeysInMap(m, allowedKeys, allowedPatterns)
		if len(invalidKeys) > 0 {
			message := fmt.Sprintf("has invalid %s: %+v", compiler.PluralProperties(len(invalidKeys)), strings.Join(invalidKeys, ", "))
			errors = append(errors, compiler.NewError(context, message))
		}
		// string name = 1;
		v1 := compiler.MapValueForKey(m, "name")
		if v1 != nil {
			x.Name, ok = v1.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for name: %+v (%T)", v1, v1)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string url = 2;
		v2 := compiler.MapValueForKey(m, "url")
		if v2 != nil {
			x.Url, ok = v2.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for url: %+v (%T)", v2, v2)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
	}
	return x, compiler.NewErrorGroupOrNil(errors)
}

// NewMaintainer creates an object of type Maintainer if possible, returning an error if not.
func NewMaintainer(in interface{}, context *compiler.Context) (*Maintainer, error) {
	errors := make([]error, 0)
	x := &Maintainer{}
	m, ok := compiler.UnpackMap(in)
	if !ok {
		message := fmt.Sprintf("has unexpected value: %+v (%T)", in, in)
		errors = append(errors, compiler.NewError(context, message))
	} else {
		allowedKeys := []string{"FN", "adr", "email", "org", "photo", "tel", "url", "vcard", "x-github", "x-twitter"}
		var allowedPatterns []*regexp.Regexp
		invalidKeys := compiler.InvalidKeysInMap(m, allowedKeys, allowedPatterns)
		if len(invalidKeys) > 0 {
			message := fmt.Sprintf("has invalid %s: %+v", compiler.PluralProperties(len(invalidKeys)), strings.Join(invalidKeys, ", "))
			errors = append(errors, compiler.NewError(context, message))
		}
		// string f_n = 1;
		v1 := compiler.MapValueForKey(m, "FN")
		if v1 != nil {
			x.Fn, ok = v1.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for FN: %+v (%T)", v1, v1)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string email = 2;
		v2 := compiler.MapValueForKey(m, "email")
		if v2 != nil {
			x.Email, ok = v2.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for email: %+v (%T)", v2, v2)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string url = 3;
		v3 := compiler.MapValueForKey(m, "url")
		if v3 != nil {
			x.Url, ok = v3.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for url: %+v (%T)", v3, v3)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string org = 4;
		v4 := compiler.MapValueForKey(m, "org")
		if v4 != nil {
			x.Org, ok = v4.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for org: %+v (%T)", v4, v4)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string adr = 5;
		v5 := compiler.MapValueForKey(m, "adr")
		if v5 != nil {
			x.Adr, ok = v5.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for adr: %+v (%T)", v5, v5)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string tel = 6;
		v6 := compiler.MapValueForKey(m, "tel")
		if v6 != nil {
			x.Tel, ok = v6.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for tel: %+v (%T)", v6, v6)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string x_twitter = 7;
		v7 := compiler.MapValueForKey(m, "x-twitter")
		if v7 != nil {
			x.XTwitter, ok = v7.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for x-twitter: %+v (%T)", v7, v7)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string x_github = 8;
		v8 := compiler.MapValueForKey(m, "x-github")
		if v8 != nil {
			x.XGithub, ok = v8.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for x-github: %+v (%T)", v8, v8)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string photo = 9;
		v9 := compiler.MapValueForKey(m, "photo")
		if v9 != nil {
			x.Photo, ok = v9.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for photo: %+v (%T)", v9, v9)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string vcard = 10;
		v10 := compiler.MapValueForKey(m, "vcard")
		if v10 != nil {
			x.Vcard, ok = v10.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for vcard: %+v (%T)", v10, v10)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
	}
	return x, compiler.NewErrorGroupOrNil(errors)
}

// NewProperty creates an object of type Property if possible, returning an error if not.
func NewProperty(in interface{}, context *compiler.Context) (*Property, error) {
	errors := make([]error, 0)
	x := &Property{}
	m, ok := compiler.UnpackMap(in)
	if !ok {
		message := fmt.Sprintf("has unexpected value: %+v (%T)", in, in)
		errors = append(errors, compiler.NewError(context, message))
	} else {
		allowedKeys := []string{"type", "url", "value"}
		var allowedPatterns []*regexp.Regexp
		invalidKeys := compiler.InvalidKeysInMap(m, allowedKeys, allowedPatterns)
		if len(invalidKeys) > 0 {
			message := fmt.Sprintf("has invalid %s: %+v", compiler.PluralProperties(len(invalidKeys)), strings.Join(invalidKeys, ", "))
			errors = append(errors, compiler.NewError(context, message))
		}
		// string type = 1;
		v1 := compiler.MapValueForKey(m, "type")
		if v1 != nil {
			x.Type, ok = v1.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for type: %+v (%T)", v1, v1)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string url = 2;
		v2 := compiler.MapValueForKey(m, "url")
		if v2 != nil {
			x.Url, ok = v2.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for url: %+v (%T)", v2, v2)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
		// string value = 3;
		v3 := compiler.MapValueForKey(m, "value")
		if v3 != nil {
			x.Value, ok = v3.(string)
			if !ok {
				message := fmt.Sprintf("has unexpected value for value: %+v (%T)", v3, v3)
				errors = append(errors, compiler.NewError(context, message))
			}
		}
	}
	return x, compiler.NewErrorGroupOrNil(errors)
}

// NewStringArray creates an object of type StringArray if possible, returning an error if not.
func NewStringArray(in interface{}, context *compiler.Context) (*StringArray, error) {
	errors := make([]error, 0)
	x := &StringArray{}
	a, ok := in.([]interface{})
	if !ok {
		message := fmt.Sprintf("has unexpected value for StringArray: %+v (%T)", in, in)
		errors = append(errors, compiler.NewError(context, message))
	} else {
		x.Value = make([]string, 0)
		for _, s := range a {
			x.Value = append(x.Value, s.(string))
		}
	}
	return x, compiler.NewErrorGroupOrNil(errors)
}

// ResolveReferences resolves references found inside Any objects.
func (m *Any) ResolveReferences(root string) (interface{}, error) {
	errors := make([]error, 0)
	return nil, compiler.NewErrorGroupOrNil(errors)
}

// ResolveReferences resolves references found inside Api objects.
func (m *Api) ResolveReferences(root string) (interface{}, error) {
	errors := make([]error, 0)
	for _, item := range m.Properties {
		if item != nil {
			_, err := item.ResolveReferences(root)
			if err != nil {
				errors = append(errors, err)
			}
		}
	}
	for _, item := range m.Contact {
		if item != nil {
			_, err := item.ResolveReferences(root)
			if err != nil {
				errors = append(errors, err)
			}
		}
	}
	return nil, compiler.NewErrorGroupOrNil(errors)
}

// ResolveReferences resolves references found inside Document objects.
func (m *Document) ResolveReferences(root string) (interface{}, error) {
	errors := make([]error, 0)
	for _, item := range m.Apis {
		if item != nil {
			_, err := item.ResolveReferences(root)
			if err != nil {
				errors = append(errors, err)
			}
		}
	}
	for _, item := range m.Include {
		if item != nil {
			_, err := item.ResolveReferences(root)
			if err != nil {
				errors = append(errors, err)
			}
		}
	}
	for _, item := range m.Maintainers {
		if item != nil {
			_, err := item.ResolveReferences(root)
			if err != nil {
				errors = append(errors, err)
			}
		}
	}
	return nil, compiler.NewErrorGroupOrNil(errors)
}

// ResolveReferences resolves references found inside Include objects.
func (m *Include) ResolveReferences(root string) (interface{}, error) {
	errors := make([]error, 0)
	return nil, compiler.NewErrorGroupOrNil(errors)
}

// ResolveReferences resolves references found inside Maintainer objects.
func (m *Maintainer) ResolveReferences(root string) (interface{}, error) {
	errors := make([]error, 0)
	return nil, compiler.NewErrorGroupOrNil(errors)
}

// ResolveReferences resolves references found inside Property objects.
func (m *Property) ResolveReferences(root string) (interface{}, error) {
	errors := make([]error, 0)
	return nil, compiler.NewErrorGroupOrNil(errors)
}

// ResolveReferences resolves references found inside StringArray objects.
func (m *StringArray) ResolveReferences(root string) (interface{}, error) {
	errors := make([]error, 0)
	return nil, compiler.NewErrorGroupOrNil(errors)
}

// ToRawInfo returns a description of Any suitable for JSON or YAML export.
func (m *Any) ToRawInfo() interface{} {
	var err error
	var info1 []yaml.MapSlice
	err = yaml.Unmarshal([]byte(m.Yaml), &info1)
	if err == nil {
		return info1
	}
	var info2 yaml.MapSlice
	err = yaml.Unmarshal([]byte(m.Yaml), &info2)
	if err == nil {
		return info2
	}
	var info3 interface{}
	err = yaml.Unmarshal([]byte(m.Yaml), &info3)
	if err == nil {
		return info3
	}
	return nil
}

// ToRawInfo returns a description of Api suitable for JSON or YAML export.
func (m *Api) ToRawInfo() interface{} {
	info := yaml.MapSlice{}
	if m.Name != "" {
		info = append(info, yaml.MapItem{Key: "name", Value: m.Name})
	}
	if m.Description != "" {
		info = append(info, yaml.MapItem{Key: "description", Value: m.Description})
	}
	if m.Image != "" {
		info = append(info, yaml.MapItem{Key: "image", Value: m.Image})
	}
	if m.HumanURL != "" {
		info = append(info, yaml.MapItem{Key: "humanURL", Value: m.HumanURL})
	}
	if m.BaseURL != "" {
		info = append(info, yaml.MapItem{Key: "baseURL", Value: m.BaseURL})
	}
	if m.Version != "" {
		info = append(info, yaml.MapItem{Key: "version", Value: m.Version})
	}
	if len(m.Tags) != 0 {
		info = append(info, yaml.MapItem{Key: "tags", Value: m.Tags})
	}
	if len(m.Properties) != 0 {
		items := make([]interface{}, 0)
		for _, item := range m.Properties {
			items = append(items, item.ToRawInfo())
		}
		info = append(info, yaml.MapItem{Key: "properties", Value: items})
	}
	// &{Name:properties Type:Property StringEnumValues:[] MapType: Repeated:true Pattern: Implicit:false Description:}
	if len(m.Contact) != 0 {
		items := make([]interface{}, 0)
		for _, item := range m.Contact {
			items = append(items, item.ToRawInfo())
		}
		info = append(info, yaml.MapItem{Key: "contact", Value: items})
	}
	// &{Name:contact Type:Maintainer StringEnumValues:[] MapType: Repeated:true Pattern: Implicit:false Description:}
	return info
}

// ToRawInfo returns a description of Document suitable for JSON or YAML export.
func (m *Document) ToRawInfo() interface{} {
	info := yaml.MapSlice{}
	if m.Name != "" {
		info = append(info, yaml.MapItem{Key: "name", Value: m.Name})
	}
	if m.Description != "" {
		info = append(info, yaml.MapItem{Key: "description", Value: m.Description})
	}
	if m.Image != "" {
		info = append(info, yaml.MapItem{Key: "image", Value: m.Image})
	}
	if m.Url != "" {
		info = append(info, yaml.MapItem{Key: "url", Value: m.Url})
	}
	if len(m.Tags) != 0 {
		info = append(info, yaml.MapItem{Key: "tags", Value: m.Tags})
	}
	if m.Created != "" {
		info = append(info, yaml.MapItem{Key: "created", Value: m.Created})
	}
	if m.Modified != "" {
		info = append(info, yaml.MapItem{Key: "modified", Value: m.Modified})
	}
	if m.SpecificationVersion != "" {
		info = append(info, yaml.MapItem{Key: "specificationVersion", Value: m.SpecificationVersion})
	}
	if len(m.Apis) != 0 {
		items := make([]interface{}, 0)
		for _, item := range m.Apis {
			items = append(items, item.ToRawInfo())
		}
		info = append(info, yaml.MapItem{Key: "apis", Value: items})
	}
	// &{Name:apis Type:Api StringEnumValues:[] MapType: Repeated:true Pattern: Implicit:false Description:}
	if len(m.Include) != 0 {
		items := make([]interface{}, 0)
		for _, item := range m.Include {
			items = append(items, item.ToRawInfo())
		}
		info = append(info, yaml.MapItem{Key: "include", Value: items})
	}
	// &{Name:include Type:Include StringEnumValues:[] MapType: Repeated:true Pattern: Implicit:false Description:}
	if len(m.Maintainers) != 0 {
		items := make([]interface{}, 0)
		for _, item := range m.Maintainers {
			items = append(items, item.ToRawInfo())
		}
		info = append(info, yaml.MapItem{Key: "maintainers", Value: items})
	}
	// &{Name:maintainers Type:Maintainer StringEnumValues:[] MapType: Repeated:true Pattern: Implicit:false Description:}
	return info
}

// ToRawInfo returns a description of Include suitable for JSON or YAML export.
func (m *Include) ToRawInfo() interface{} {
	info := yaml.MapSlice{}
	if m.Name != "" {
		info = append(info, yaml.MapItem{Key: "name", Value: m.Name})
	}
	if m.Url != "" {
		info = append(info, yaml.MapItem{Key: "url", Value: m.Url})
	}
	return info
}

// ToRawInfo returns a description of Maintainer suitable for JSON or YAML export.
func (m *Maintainer) ToRawInfo() interface{} {
	info := yaml.MapSlice{}
	if m.Fn != "" {
		info = append(info, yaml.MapItem{Key: "FN", Value: m.Fn})
	}
	if m.Email != "" {
		info = append(info, yaml.MapItem{Key: "email", Value: m.Email})
	}
	if m.Url != "" {
		info = append(info, yaml.MapItem{Key: "url", Value: m.Url})
	}
	if m.Org != "" {
		info = append(info, yaml.MapItem{Key: "org", Value: m.Org})
	}
	if m.Adr != "" {
		info = append(info, yaml.MapItem{Key: "adr", Value: m.Adr})
	}
	if m.Tel != "" {
		info = append(info, yaml.MapItem{Key: "tel", Value: m.Tel})
	}
	if m.XTwitter != "" {
		info = append(info, yaml.MapItem{Key: "x-twitter", Value: m.XTwitter})
	}
	if m.XGithub != "" {
		info = append(info, yaml.MapItem{Key: "x-github", Value: m.XGithub})
	}
	if m.Photo != "" {
		info = append(info, yaml.MapItem{Key: "photo", Value: m.Photo})
	}
	if m.Vcard != "" {
		info = append(info, yaml.MapItem{Key: "vcard", Value: m.Vcard})
	}
	return info
}

// ToRawInfo returns a description of Property suitable for JSON or YAML export.
func (m *Property) ToRawInfo() interface{} {
	info := yaml.MapSlice{}
	if m.Type != "" {
		info = append(info, yaml.MapItem{Key: "type", Value: m.Type})
	}
	if m.Url != "" {
		info = append(info, yaml.MapItem{Key: "url", Value: m.Url})
	}
	if m.Value != "" {
		info = append(info, yaml.MapItem{Key: "value", Value: m.Value})
	}
	return info
}

// ToRawInfo returns a description of StringArray suitable for JSON or YAML export.
func (m *StringArray) ToRawInfo() interface{} {
	return m.Value
}

var ()
