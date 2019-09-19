package an

import (
	"log"

	"bitbucket.org/Axxonsoft/bl/config"
	"bitbucket.org/Axxonsoft/bl/primitive"
)

// PropertiesBuilder utility class for composing unit properties
type PropertiesBuilder struct {
	properties []*config.Property
}

// MakePropBuilder returns new PropertiesBuilder
func MakePropBuilder() *PropertiesBuilder {
	return &PropertiesBuilder{}
}

func (pb *PropertiesBuilder) ensure() *PropertiesBuilder {
	if pb == nil {
		return MakePropBuilder()
	}
	return pb
}

// AddString adds string property
func (pb *PropertiesBuilder) AddString(id, value string) *PropertiesBuilder {
	pb = pb.ensure()
	pb.properties = append(pb.properties,
		&config.Property{Id: id, Value: &config.Property_ValueString{value}})
	return pb
}

// AddProp adds property of any type using reflection
func (pb *PropertiesBuilder) AddProp(id string, value interface{}) *PropertiesBuilder {
	switch v := value.(type) {
	case int:
		return pb.AddInt(id, v)
	case float64:
		return pb.AddDouble(id, v)
	case string:
		return pb.AddString(id, v)
	case bool:
		return pb.AddBool(id, v)
	case *primitive.Polyline:
		return pb.AddPolyline(id, v)
	case *primitive.HsvColorRange:
		return pb.AddHsvColorRange(id, v)
	case *primitive.Rectangle:
		pb.AddRectangle(id, v)
	default:
		log.Panicf("unknown type %T of value: %+v for id %s", value, value, id)
	}
	return pb
}

// AddInt adds int property
func (pb *PropertiesBuilder) AddInt(id string, value int) *PropertiesBuilder {
	pb = pb.ensure()
	pb.properties = append(pb.properties,
		&config.Property{Id: id, Value: &config.Property_ValueInt32{ValueInt32: int32(value)}})
	return pb
}

// AddDouble adds double property
func (pb *PropertiesBuilder) AddDouble(id string, value float64) *PropertiesBuilder {
	pb = pb.ensure()
	pb.properties = append(pb.properties, &config.Property{Id: id, Value: &config.Property_ValueDouble{ValueDouble: value}})
	return pb
}

// AddHsvColorRange HsvColorRange int property
func (pb *PropertiesBuilder) AddHsvColorRange(id string, value *primitive.HsvColorRange) *PropertiesBuilder {
	pb = pb.ensure()
	pb.properties = append(pb.properties, &config.Property{Id: id, Value: &config.Property_ValueHsvColorRange{ValueHsvColorRange: value}})
	return pb
}

// AddBool adds bool property
func (pb *PropertiesBuilder) AddBool(id string, value bool) *PropertiesBuilder {
	pb = pb.ensure()
	pb.properties = append(pb.properties,
		&config.Property{Id: id, Value: &config.Property_ValueBool{ValueBool: value}})
	return pb
}

// AddPolyline adds polyline property
func (pb *PropertiesBuilder) AddPolyline(id string, value *primitive.Polyline) *PropertiesBuilder {
	pb = pb.ensure()
	pb.properties = append(pb.properties,
		&config.Property{Id: id, Value: &config.Property_ValuePolyline{value}})
	return pb
}

// AddRectangle adds rectangle property
func (pb *PropertiesBuilder) AddRectangle(id string, value *primitive.Rectangle) *PropertiesBuilder {
	pb = pb.ensure()
	pb.properties = append(pb.properties,
		&config.Property{Id: id, Value: &config.Property_ValueRectangle{value}})
	return pb
}

// Get returns all properties
func (pb *PropertiesBuilder) Get() []*config.Property {
	if pb == nil {
		return nil
	}
	return pb.properties
}

// Empty returns is empty properties
func (pb *PropertiesBuilder) Empty() bool {
	if pb == nil {
		return true
	}
	return len(pb.properties) == 0
}
