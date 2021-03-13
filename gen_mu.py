#!/usr/bin/env python
"""
Generates sum, min, max, abs functions for Go numeric types.
"""
# pylint: disable=R,C
typesfx = (
    ("int", ""),
    ("int8", "I8"),
    ("int16", "I16"),
    ("int32", "I32"),
    ("int64", "I64"),
    ("uint8", "U8"),
    ("uint16", "U16"),
    ("uint32", "U32"),
    ("uint64", "U64"),
    ("float32", "F32"),
    ("float64", "F64"),
)


def sums(outf):
    "Variadic sum funcs for all types"
    tp = """
    // {fname} returns the sum of its arguments.
    // An empty argument list raises a panic.
    func {fname}(nums ...{T}) {T} {{
    if len(nums) == 0 {{
        panic("{fname} argument list empty")
        }}
    var total {T} = 0
    for _, num := range nums {{
        total += num
    }}
    return total
    }}
    """
    for t, sfx in typesfx:
        d = dict(fname="Sum" + sfx, T=t)
        print(tp.format(**d), file=outf)


def sumtests(outf):
    tp = """
    func Test{fname}(t *testing.T) {{
        if val := {fname}({args}); val != {want} {{
            t.Errorf("Expected %v, got %v.", val, {want})
        }}
    }}
    """
    for t, sfx in typesfx:
        if not t.startswith("u"):  # unsigned type
            d = dict(fname="Sum" + sfx, args="1, -2, 3", want=2)
        else:
            d = dict(fname="Sum" + sfx, args="1, 2, 3", want=6)
        print(tp.format(**d), file=outf)


def mins(outf):
    "Variadic min funcs for all types"
    tp = """
    // {fname} returns the minimum of its arguments.
    // An empty argument list raises a panic.
    func {fname}(nums ...{T}) {T} {{
    if len(nums) == 0 {{
        panic("{fname} argument list empty")
        }}
    min := nums[0]
    for _, num := range nums {{
        if num < min {{
           min = num
           }}
    }}
    return min
    }}
    """
    for t, sfx in typesfx:
        d = dict(fname="Min" + sfx, T=t)
        print(tp.format(**d), file=outf)


def mintests(outf):
    tp = """
    func Test{fname}(t *testing.T) {{
        if val := {fname}({args}); val != {want} {{
            t.Errorf("Expected %v, got %v.", val, {want})
        }}
    }}
    """
    for t, sfx in typesfx:
        if not t.startswith("u"):  # unsigned type
            d = dict(fname="Min" + sfx, args="1, -2, 3", want=-2)
        else:
            d = dict(fname="Min" + sfx, args="1, 2, 3", want=1)
        print(tp.format(**d), file=outf)


def maxes(outf):
    "Variadic min funcs for all types"
    tp = """
    // {fname} returns the maximum of its arguments.
    // An empty argument list raises a panic.
    func {fname}(nums ...{T}) {T} {{
    if len(nums) == 0 {{
        panic("{fname} argument list empty")
        }}
    max := nums[0]
    for _, num := range nums {{
        if num > max {{
           max = num
           }}
    }}
    return max
    }}
    """
    for t, sfx in typesfx:
        d = dict(fname="Max" + sfx, T=t)
        print(tp.format(**d), file=outf)


def maxtests(outf):
    tp = """
    func Test{fname}(t *testing.T) {{
        if val := {fname}({args}); val != {want} {{
            t.Errorf("Expected %v, got %v.", val, {want})
        }}
    }}
    """
    for t, sfx in typesfx:
        if not t.startswith("u"):  # unsigned type
            d = dict(fname="Max" + sfx, args="1, -2, -3", want=1)
        else:
            d = dict(fname="Max" + sfx, args="1, 2, 3", want=3)
        print(tp.format(**d), file=outf)


def abses(outf):
    "monadic abs funcs for all types"
    tpm = """
    // {fname} returns the absolute values of its argument.
    func {fname}(num {T}) {T} {{
    var abs {T}
        switch num < 0 {{
        case true:
           abs = -num
        default:
           abs = num
        }}
    return abs
    }}

    """
    # Variadic abs funcs for all types
    tpv = """
    // {fname} returns the absolute values of its arguments.
    // An empty argument list raises a panic.
    func {fname}(nums ...{T}) []{T} {{
    if len(nums) == 0 {{
        panic("{fname} argument list empty")
        }}
    var abs []{T}
    for _, num := range nums {{
        switch num < 0 {{
        case true:
           abs = append(abs, -num)
        default:
           abs = append(abs, num)
        }}
    }}
    return abs
    }}
    """
    for t, sfx in typesfx:
        if t.startswith("u"):
            continue  ## don't need abs for unsigned types
        d = dict(fname="Abs" + sfx, T=t)
        print(tpm.format(**d), file=outf)
        d = dict(fname="VAbs" + sfx, T=t)
        print(tpv.format(**d), file=outf)


def abstests(outf):
    tpm = """
    func Test{fname}(t *testing.T) {{
        if val := {fname}({args}); !(val == {want}) {{
            t.Errorf("Expected %v, got %v.", val, {want})
        }}
    }}
    """
    tpv = """
    func Test{fname}(t *testing.T) {{
        if val := {fname}({args}); !reflect.DeepEqual(val, {want}) {{
            t.Errorf("Expected %v, got %v.", val, {want})
        }}
    }}
    """
    for t, sfx in typesfx:
        if not t.startswith("u"):  # unsigned type
            ## monadic abses
            d = dict(fname="Abs" + sfx, args=-1, want=1)
            print(tpm.format(**d), file=outf)
            ## variadic abses
            want = "[]{}{{1, 2, 3}}".format(t)
            d = dict(fname="VAbs" + sfx, args="1, -2, -3", want=want)
            print(tpv.format(**d), file=outf)


def comptests(outf):
    t = """
    func TestCloseEnough(t *testing.T) {
        if !CloseEnough(1.0, 1.001, .001) {
            t.Errorf("Expected 1.0 and 1.001 to be close enough with tolerance .001")
        }
        if CloseEnough(1.0, 1.001, .0001) {
            t.Errorf("Expected 1.0 and 1.001 not to be close enough with tolerance .0001")
        }
    }
    """
    print(t, file=outf)


## Write functions
with open("mu_g.go", "w+") as f:
    print("// Code generated by gen_mu.py. DO NOT EDIT.\n", file=f)
    print("package mu\n", file=f)
    sums(f)
    mins(f)
    maxes(f)
    abses(f)

## Write tests
with open("mu_test.go", "w+") as f:
    print("// Code generated by gen_mu.py. DO NOT EDIT.\n", file=f)
    print("package mu\n", file=f)
    print('import "reflect"', file=f)
    print('import "testing"\n', file=f)
    sumtests(f)
    mintests(f)
    maxtests(f)
    abstests(f)
    comptests(f)
