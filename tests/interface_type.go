package tests

// easyjson:json
type WrapperType struct {
	Inner any `json:"Inner"`
}

// easyjson:json
type InterfaceType struct {
	Field1 int `json:"Field1"`
}

// type EmbeddedInnerType2 struct {
// 	Field3 int
// }

// var embeddedTypeValue EmbeddedType

// func init() {
// 	embeddedTypeValue.Field1 = 1
// 	embeddedTypeValue.Field2 = 2
// 	embeddedTypeValue.Inner.Field1 = 3
// 	embeddedTypeValue.Field3 = 4
// }

// var embeddedTypeValueString = `{"Inner":{"Field1":3},"Field2":2,"named":{"Field3":4},"Field1":1}`
