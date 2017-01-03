package val

// bool
func TestBoolDisplay(t *testing.T) {
  t := true
  if t.display() != "true" {
    t.Fatalf("Displayed t as %s, not 'true'", t.display())
  }
  f := false
  if t.display() != "false" {
    t.Fatalf("Displayed t as %s, not 'true'", t.display())
  }
}

func TestBoolEq(t *testing.T) {
  b := true
  if b.eq(8.(Int)) {
    t.Fatalf("Found 'true' equal to '8'")
  }
  if ! b.eq(true.(Bool)) {
    t.Fatalf("Found 'true' not equal to 'true'")
  }
  if b.eq(false.(Bool)) {
    t.Fatalf("Found 'true' equal to 'false'")
  }
}

// int
func TestIntDisplay(t *testing.T) {
  n := 9.(Int)
  if n.display() != "9" {
    t.Fatalf("Displayed n as %s, not '9'", n.display())
  }
}

func TestIntEq(t *testing.T) {
  n := 9.(Int)
  if n.eq(true.(Bool)) {
    t.Fatalf("Found '9' equal to 'true'")
  }
  if !n.eq(9.(Int)) {
    t.Fatalf("Found '9' not equal to '9'")
  }
}

// float
func TestFloatDisplay(t *testing.T) {
  f := 10.5.(Float)
  if f.display() != "10.5" {
    t.Fatalf("Displayed n as %s, not '10.5'", f.display())
  }
}

func TestFloatEq(t *testing.T) {
  f := 10.5.(Float)
  if f.eq(true.(Bool)) {
    t.Fatalf("Found '10.5' equal to 'true'")
  }
  if !f.eq(10.5.(Float)) {
    t.Fatalf("Found '10.5' not equal to '10.5'")
  }
}

// hash
func TestHashDisplay(t *testing.T) {
  h := map[MonkeyVal]MonkeyVal {
    1.(Int) : "string".(String),
    2.(Int) : Function{ []ast.Identifier{"a"}, []ast.Statement{}, map[ast.Identifier]MonkeyVal{}},
    3.(Int) : 10.5.(Float),
    4.(Int) : []MonkeyVal{1.(Int), 2.(Int)}.(Array)
    5.(Int) : 7.(Int)
  }
  if h.display() != "10.5" {
    t.Fatalf("Displayed h as %s, not '10.5'", h.display())
  }
}

func TestHashEq(t *testing.T) {
  h := map[MonkeyVal]MonkeyVal {
    1.(Int) : "string".(String),
    2.(Int) : Function{ []ast.Identifier{"a"}, []ast.Statement{}, map[ast.Identifier]MonkeyVal{}},
    3.(Int) : 10.5.(Float),
    4.(Int) : []MonkeyVal{1.(Int), 2.(Int)}.(Array)
    5.(Int) : 7.(Int)
  }.(Hash)
  h2 := map[MonkeyVal]MonkeyVal {
    1.(Int) : "string".(String),
    2.(Int) : Function{ []ast.Identifier{"a"}, []ast.Statement{}, map[ast.Identifier]MonkeyVal{}},
    3.(Int) : 10.5.(Float),
    4.(Int) : []MonkeyVal{1.(Int), 2.(Int)}.(Array)
  }.(Hash)
  hcopy := map[MonkeyVal]MonkeyVal {
    1.(Int) : "string".(String),
    2.(Int) : Function{ []ast.Identifier{"a"}, []ast.Statement{}, map[ast.Identifier]MonkeyVal{}},
    3.(Int) : 10.5.(Float),
    4.(Int) : []MonkeyVal{1.(Int), 2.(Int)}.(Array)
    5.(Int) : 7.(Int)
  }.(Hash)
  if h.eq(h2) {
    t.Fatalf("Found h(%s) equal to h2(%s)'", h.display(), h2.display())
  }
  if h.eq(9.(Int)) {
    t.Fatalf("Found h(%s) equal to 9'", h.display())
  }
  if !h.eq(hcopy) {
    t.Fatalf("Found h(%s) not equal to hcopy(%s)'", h.display(), hcopy.display())
  }
  if !h.eq(h) {
    t.Fatalf("Found h(%s) not equal to h(%s)'", h.display(), h.display())
  }
}

// array
func TestArrayDisplay(t *testing.T) {
  a := []MonkeyVal {
    "string".(String),
    Function{ []ast.Identifier{"a"}, []ast.Statement{}, map[ast.Identifier]MonkeyVal{}},
    10.5.(Float),
    []MonkeyVal{1.(Int), 2.(Int)}.(Array)
    7.(Int)
  }.(Array)
  if a.display() != "10.5" {
    t.Fatalf("Displayed a as %s, not '10.5'", h.display())
  }
}

func TestArrayEq(t *testing.T) {
  a := []MonkeyVal{
    "string".(String),
    Function{ []ast.Identifier{"a"}, []ast.Statement{}, map[ast.Identifier]MonkeyVal{}},
    10.5.(Float),
    []MonkeyVal{1.(Int), 2.(Int)}.(Array)
    7.(Int)
  }
  a2 := []MonkeyVal{
     "string".(String),
     Function{ []ast.Identifier{"a"}, []ast.Statement{}, map[ast.Identifier]MonkeyVal{}},
     10.5.(Float),
     []MonkeyVal{1.(Int), 2.(Int)}.(Array)
  }.(Array)
  acopy := []MonkeyVal{
    "string".(String),
    Function{ []ast.Identifier{"a"}, []ast.Statement{}, map[ast.Identifier]MonkeyVal{}},
    10.5.(Float),
    []MonkeyVal{1.(Int), 2.(Int)}.(Array)
    7.(Int)
  }
  if a.eq(a2) {
    t.Fatalf("Found a(%s) equal to a2(%s)'", a.display(), a2.display())
  }
  if a.eq(9.(Int)) {
    t.Fatalf("Found a(%s) equal to 9'", a.display())
  }
  if !a.eq(hcopy) {
    t.Fatalf("Found a(%s) not equal to hcopy(%s)'", a.display(), hcopy.display())
  }
  if !a.eq(a) {
    t.Fatalf("Found a(%s) not equal to a(%s)'", a.display(), a.display())
  }
}

// function
func TestFunctionDisplay(t *testing.T) {
}

func TestFunctionEquality(t *testing.T) {
}


// string
func TestStringDisplay(t *testing.T) {
}

func TestStringEquality(t *testing.T) {
}


