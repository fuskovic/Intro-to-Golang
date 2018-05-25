package main

import (
	"fmt"
	"reflect"
	//have to use reflect to access back tic tags
)

//******************************************* Compositions and embedding plus adding tags to structs

type Animal struct {
	//adding tags to structs
	Name string `required max:"100"`
	//name field is required and max length of value is 100 in the even of filling out a web form or something
	Origin string
}

type Bird struct {
	//embed animal struct to use composition to inherit Name and Origin above
	Animal
	//now Bird types can use Name and Origin
	SpeedKPH float32
	CanFly   bool
}

func main() {

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)

	/*
		NOTE: a Bird is its own struct, it is not type Animal
		its just composing its field values through embedding.

		Think of this as " a bird has animal like characteristics"
		But is not the same thing as an animal even though in real
		life birds are totally animals.
	*/

	//lets illustrate the above through setting Animal struct field values

	c := Bird{
		/*
			Keep in mind Go doesn't like the word inherited
			becuase the Animal and Bird structs are still technically
			independent from one another.

			Notice below, an instance of animal is created inside this new Bird
			Struct(c) and its field values are then set.

		*/
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(c.Name)

	//Adding tags to structs
	t := reflect.TypeOf(Animal{})
	//Typeof needs the struct that containst the tag
	field, _ := t.FieldByName("Name")
	/*
		Write only is set to throw away the bool type returned from FieldByName

		FieldByName takes the field name( in this case "Name")
	*/
	fmt.Println(field.Tag)
}
