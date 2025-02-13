new things learned from each project:
* reflection
    * reflect = to get data type property in execution time.
    * useful for utilizing fields name / type in struct (`reflect.ValueOf(x).Type().Field(0).Name`)

* setValues
    * if `a = struct {X string, Y int, Z float64}`
        * `r := reflect.ValueOf(a)`
        
        * `r.Type().Field(0).Name` => X as plain key string
        * `r.Type().Field(0).Type().Kind()` to make a comparable type (`typeOfX == reflect.int`)

        * `r.Field(0).Interface()` => current value of X

        * `r.Field(0).Name` DOESN'T EXIST, so must go through `r.Type().Field(0).Name`

        * `r.Field(0).SetString("new val")` to set struct field (X)

    * init with `r := reflect.ValueOf(&A).Elem()` to make `A` modifiable

* methods
    * struct can have methods to modify itself
        * `struct s`
            * declare `func (var1 *s) modifyField (b int) => var1.field = b`
            * then `s1.modifyField(42); (s1.field == 42) == true`

* sort
    * interface == set of requirements to be implemented by data type
    * `sort.Sort` accepts `Interface` as param
        * `sort.Sort(S2Slice{})` == `S2Slice` must implement `Interface` requirement from `sort` package, such as: `Len()`, `Less()`, `Swap()` 

* empty
    * a function can accept empty interface as param.
    * maybe similar as `any` from typescript, sort of.

* typeSwitch
    * apply `.(type)` to `interface{}` as param to check data type and act accordingly.