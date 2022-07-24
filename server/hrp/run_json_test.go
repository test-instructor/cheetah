package hrp

import (
	"testing"
)

func TestJsonRunner(t *testing.T) {
	jsonString := `{
    "config": {
        "name": "demo with complex mechanisms",
        "base_url": "https://postman-echo.com",
        "variables": {
            "a": "${sum(10, 2.3)}",
            "b": 3.45,
            "n": "${sum_ints(1, 2, 2)}",
            "varFoo1": "${gen_random_string($n)}",
            "varFoo2": "${max($a, $b)}"
        }
    },
    "teststeps": [
        {
            "name": "transaction 1 start",
            "transaction": {
                "name": "tran1",
                "type": "start"
            }
        },
        {
            "name": "get with params",
            "request": {
                "method": "GET",
                "url": "/get",
                "params": {
                    "foo1": "$varFoo1",
                    "foo2": "$varFoo2"
                },
                "headers": {
                    "User-Agent": "HttpRunnerPlus"
                }
            },
            "variables": {
                "b": 34.5,
                "n": 3,
                "name": "get with params",
                "varFoo2": "${max($a, $b)}"
            },
            "setup_hooks": [
                "${setup_hook_example($name)}"
            ],
            "teardown_hooks": [
                "${teardown_hook_example($name)}"
            ],
            "extract": {
                "varFoo1": "body.args.foo1"
            },
            "validate": [
                {
                    "check": "status_code",
                    "assert": "equals",
                    "expect": 200,
                    "msg": "check response status code"
                },
                {
                    "check": "headers.\"Content-Type\"",
                    "assert": "startswith",
                    "expect": "application/json"
                },
                {
                    "check": "body.args.foo1",
                    "assert": "length_equals",
                    "expect": 5,
                    "msg": "check args foo1"
                },
                {
                    "check": "$varFoo1",
                    "assert": "length_equals",
                    "expect": 5,
                    "msg": "check args foo1"
                },
                {
                    "check": "body.args.foo2",
                    "assert": "equals",
                    "expect": "34.5",
                    "msg": "check args foo2"
                }
            ]
        },
        {
            "name": "transaction 1 end",
            "transaction": {
                "name": "tran1",
                "type": "end"
            }
        },
        {
            "name": "post json data",
            "request": {
                "method": "POST",
                "url": "/post",
                "body": {
                    "foo1": "$varFoo1",
                    "foo2": "${max($a, $b)}"
                }
            },
            "validate": [
                {
                    "check": "status_code",
                    "assert": "equals",
                    "expect": 200,
                    "msg": "check status code"
                },
                {
                    "check": "body.json.foo1",
                    "assert": "length_equals",
                    "expect": 5,
                    "msg": "check args foo1"
                },
                {
                    "check": "body.json.foo2",
                    "assert": "equals",
                    "expect": 12.3,
                    "msg": "check args foo2"
                }
            ]
        },
        {
            "name": "post form data",
            "request": {
                "method": "POST",
                "url": "/post",
                "headers": {
                    "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8"
                },
                "body": {
                    "foo1": "$varFoo1",
                    "foo2": "${max($a, $b)}",
                    "time": "${get_timestamp()}"
                }
            },
            "extract": {
                "varTime": "body.form.time"
            },
            "validate": [
                {
                    "check": "status_code",
                    "assert": "equals",
                    "expect": 200,
                    "msg": "check status code"
                },
                {
                    "check": "body.form.foo1",
                    "assert": "length_equals",
                    "expect": 5,
                    "msg": "check args foo1"
                },
                {
                    "check": "body.form.foo2",
                    "assert": "equals",
                    "expect": "12.3",
                    "msg": "check args foo2"
                }
            ]
        },
        {
            "name": "get with timestamp",
            "request": {
                "method": "GET",
                "url": "/get",
                "params": {
                    "time": "$varTime"
                }
            },
            "validate": [
                {
                    "check": "body.args.time",
                    "assert": "length_equals",
                    "expect": 13,
                    "msg": "check extracted var timestamp"
                }
            ]
        }
    ]
}`
	BuildHashicorpGoPlugin()
	defer RemoveHashicorpGoPlugin()
	testcase3 := &TestCaseJson{jsonString, 1}
	testCase, _ := testcase3.ToTestCase()
	err := NewRunner(t).Run(testCase)
	if err != nil {
		t.Fatalf("run testcase error: %v", err)
	}
}
