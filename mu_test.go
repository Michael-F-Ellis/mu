// Code generated by gen_mu.py. DO NOT EDIT.

package mu

import "reflect"
import "testing"

func TestSum(t *testing.T) {
	if val := Sum(1, -2, 3); val != 2 {
		t.Errorf("Expected %v, got %v.", val, 2)
	}
}

func TestSumI8(t *testing.T) {
	if val := SumI8(1, -2, 3); val != 2 {
		t.Errorf("Expected %v, got %v.", val, 2)
	}
}

func TestSumI16(t *testing.T) {
	if val := SumI16(1, -2, 3); val != 2 {
		t.Errorf("Expected %v, got %v.", val, 2)
	}
}

func TestSumI32(t *testing.T) {
	if val := SumI32(1, -2, 3); val != 2 {
		t.Errorf("Expected %v, got %v.", val, 2)
	}
}

func TestSumI64(t *testing.T) {
	if val := SumI64(1, -2, 3); val != 2 {
		t.Errorf("Expected %v, got %v.", val, 2)
	}
}

func TestSumU8(t *testing.T) {
	if val := SumU8(1, 2, 3); val != 6 {
		t.Errorf("Expected %v, got %v.", val, 6)
	}
}

func TestSumU16(t *testing.T) {
	if val := SumU16(1, 2, 3); val != 6 {
		t.Errorf("Expected %v, got %v.", val, 6)
	}
}

func TestSumU32(t *testing.T) {
	if val := SumU32(1, 2, 3); val != 6 {
		t.Errorf("Expected %v, got %v.", val, 6)
	}
}

func TestSumU64(t *testing.T) {
	if val := SumU64(1, 2, 3); val != 6 {
		t.Errorf("Expected %v, got %v.", val, 6)
	}
}

func TestSumF32(t *testing.T) {
	if val := SumF32(1, -2, 3); val != 2 {
		t.Errorf("Expected %v, got %v.", val, 2)
	}
}

func TestSumF64(t *testing.T) {
	if val := SumF64(1, -2, 3); val != 2 {
		t.Errorf("Expected %v, got %v.", val, 2)
	}
}

func TestMin(t *testing.T) {
	if val := Min(1, -2, 3); val != -2 {
		t.Errorf("Expected %v, got %v.", val, -2)
	}
}

func TestMinI8(t *testing.T) {
	if val := MinI8(1, -2, 3); val != -2 {
		t.Errorf("Expected %v, got %v.", val, -2)
	}
}

func TestMinI16(t *testing.T) {
	if val := MinI16(1, -2, 3); val != -2 {
		t.Errorf("Expected %v, got %v.", val, -2)
	}
}

func TestMinI32(t *testing.T) {
	if val := MinI32(1, -2, 3); val != -2 {
		t.Errorf("Expected %v, got %v.", val, -2)
	}
}

func TestMinI64(t *testing.T) {
	if val := MinI64(1, -2, 3); val != -2 {
		t.Errorf("Expected %v, got %v.", val, -2)
	}
}

func TestMinU8(t *testing.T) {
	if val := MinU8(1, 2, 3); val != 1 {
		t.Errorf("Expected %v, got %v.", val, 1)
	}
}

func TestMinU16(t *testing.T) {
	if val := MinU16(1, 2, 3); val != 1 {
		t.Errorf("Expected %v, got %v.", val, 1)
	}
}

func TestMinU32(t *testing.T) {
	if val := MinU32(1, 2, 3); val != 1 {
		t.Errorf("Expected %v, got %v.", val, 1)
	}
}

func TestMinU64(t *testing.T) {
	if val := MinU64(1, 2, 3); val != 1 {
		t.Errorf("Expected %v, got %v.", val, 1)
	}
}

func TestMinF32(t *testing.T) {
	if val := MinF32(1, -2, 3); val != -2 {
		t.Errorf("Expected %v, got %v.", val, -2)
	}
}

func TestMinF64(t *testing.T) {
	if val := MinF64(1, -2, 3); val != -2 {
		t.Errorf("Expected %v, got %v.", val, -2)
	}
}

func TestMax(t *testing.T) {
	if val := Max(1, -2, -3); val != 1 {
		t.Errorf("Expected %v, got %v.", val, 1)
	}
}

func TestMaxI8(t *testing.T) {
	if val := MaxI8(1, -2, -3); val != 1 {
		t.Errorf("Expected %v, got %v.", val, 1)
	}
}

func TestMaxI16(t *testing.T) {
	if val := MaxI16(1, -2, -3); val != 1 {
		t.Errorf("Expected %v, got %v.", val, 1)
	}
}

func TestMaxI32(t *testing.T) {
	if val := MaxI32(1, -2, -3); val != 1 {
		t.Errorf("Expected %v, got %v.", val, 1)
	}
}

func TestMaxI64(t *testing.T) {
	if val := MaxI64(1, -2, -3); val != 1 {
		t.Errorf("Expected %v, got %v.", val, 1)
	}
}

func TestMaxU8(t *testing.T) {
	if val := MaxU8(1, 2, 3); val != 3 {
		t.Errorf("Expected %v, got %v.", val, 3)
	}
}

func TestMaxU16(t *testing.T) {
	if val := MaxU16(1, 2, 3); val != 3 {
		t.Errorf("Expected %v, got %v.", val, 3)
	}
}

func TestMaxU32(t *testing.T) {
	if val := MaxU32(1, 2, 3); val != 3 {
		t.Errorf("Expected %v, got %v.", val, 3)
	}
}

func TestMaxU64(t *testing.T) {
	if val := MaxU64(1, 2, 3); val != 3 {
		t.Errorf("Expected %v, got %v.", val, 3)
	}
}

func TestMaxF32(t *testing.T) {
	if val := MaxF32(1, -2, -3); val != 1 {
		t.Errorf("Expected %v, got %v.", val, 1)
	}
}

func TestMaxF64(t *testing.T) {
	if val := MaxF64(1, -2, -3); val != 1 {
		t.Errorf("Expected %v, got %v.", val, 1)
	}
}

func TestAbs(t *testing.T) {
	if val := Abs(-1); !(val == 1) {
		t.Errorf("Expected %v, got %v.", val, 1)
	}
}

func TestVAbs(t *testing.T) {
	if val := VAbs(1, -2, -3); !reflect.DeepEqual(val, []int{1, 2, 3}) {
		t.Errorf("Expected %v, got %v.", val, []int{1, 2, 3})
	}
}

func TestAbsI8(t *testing.T) {
	if val := AbsI8(-1); !(val == 1) {
		t.Errorf("Expected %v, got %v.", val, 1)
	}
}

func TestVAbsI8(t *testing.T) {
	if val := VAbsI8(1, -2, -3); !reflect.DeepEqual(val, []int8{1, 2, 3}) {
		t.Errorf("Expected %v, got %v.", val, []int8{1, 2, 3})
	}
}

func TestAbsI16(t *testing.T) {
	if val := AbsI16(-1); !(val == 1) {
		t.Errorf("Expected %v, got %v.", val, 1)
	}
}

func TestVAbsI16(t *testing.T) {
	if val := VAbsI16(1, -2, -3); !reflect.DeepEqual(val, []int16{1, 2, 3}) {
		t.Errorf("Expected %v, got %v.", val, []int16{1, 2, 3})
	}
}

func TestAbsI32(t *testing.T) {
	if val := AbsI32(-1); !(val == 1) {
		t.Errorf("Expected %v, got %v.", val, 1)
	}
}

func TestVAbsI32(t *testing.T) {
	if val := VAbsI32(1, -2, -3); !reflect.DeepEqual(val, []int32{1, 2, 3}) {
		t.Errorf("Expected %v, got %v.", val, []int32{1, 2, 3})
	}
}

func TestAbsI64(t *testing.T) {
	if val := AbsI64(-1); !(val == 1) {
		t.Errorf("Expected %v, got %v.", val, 1)
	}
}

func TestVAbsI64(t *testing.T) {
	if val := VAbsI64(1, -2, -3); !reflect.DeepEqual(val, []int64{1, 2, 3}) {
		t.Errorf("Expected %v, got %v.", val, []int64{1, 2, 3})
	}
}

func TestAbsF32(t *testing.T) {
	if val := AbsF32(-1); !(val == 1) {
		t.Errorf("Expected %v, got %v.", val, 1)
	}
}

func TestVAbsF32(t *testing.T) {
	if val := VAbsF32(1, -2, -3); !reflect.DeepEqual(val, []float32{1, 2, 3}) {
		t.Errorf("Expected %v, got %v.", val, []float32{1, 2, 3})
	}
}

func TestAbsF64(t *testing.T) {
	if val := AbsF64(-1); !(val == 1) {
		t.Errorf("Expected %v, got %v.", val, 1)
	}
}

func TestVAbsF64(t *testing.T) {
	if val := VAbsF64(1, -2, -3); !reflect.DeepEqual(val, []float64{1, 2, 3}) {
		t.Errorf("Expected %v, got %v.", val, []float64{1, 2, 3})
	}
}

func TestCloseEnough(t *testing.T) {
	if !CloseEnough(1.0, 1.001, .001) {
		t.Errorf("Expected 1.0 and 1.001 to be close enough with tolerance .001")
	}
	if CloseEnough(1.0, 1.001, .0001) {
		t.Errorf("Expected 1.0 and 1.001 not to be close enough with tolerance .0001")
	}
}