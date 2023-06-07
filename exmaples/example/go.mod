
module icode.baidu.com/baidu/bce-soe/rule-engine/examples/example

go 1.19

 // use replace directive for module local test

require (
	icode.baidu.com/baidu/bce-soe/rule-engine v0.0.0
)
replace (
	icode.baidu.com/baidu/bce-soe/rule-engine v0.0.0 => ../../
)