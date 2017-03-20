package querybuilder

type opType string

const (
	opType_logical    opType = "logical"
	opType_comparison opType = "comparison"
	opType_unknown    opType = "unknown"
	opType_NotComp    opType = "notComp"
)

const (
	operatorMongo_gt  = "$gt"
	operatorMongo_lt  = "$lt"
	operatorMongo_gte = "$gte"
	operatorMongo_lte = "$lte"
	operatorMongo_eq  = "$eq"
	operatorMongo_ne  = "$ne"

	operatorMongo_and = "$and"
	operatorMongo_or  = "$or"
	operatorMongo_not = "$not"
)

const (
	operator_startBracket = "("
	operator_CloseBracket = ")"
	operator_gt           = ">"
	operator_gte          = ">="
	operator_lt           = "<"
	operator_lte          = "<="
	operator_eq           = "="
	operator_ne           = "!="
	operator_and          = "&&"
	operator_or           = "||"
	operator_not          = "!"

	operator_not_gt  = "!>"
	operator_not_gte = "!>="
	operator_not_lt  = "!<"
	operator_not_lte = "!<="
)

var comOpConvMap map[string]string

func init() {
	comOpConvMap = make(map[string]string)
	comOpConvMap[operator_gt] = operatorMongo_gt
	comOpConvMap[operator_lt] = operatorMongo_lt
	comOpConvMap[operator_gte] = operatorMongo_gte
	comOpConvMap[operator_lte] = operatorMongo_lte
	comOpConvMap[operator_eq] = operatorMongo_eq
	comOpConvMap[operator_ne] = operatorMongo_ne

	comOpConvMap[operator_and] = operatorMongo_and
	comOpConvMap[operator_or] = operatorMongo_or
	comOpConvMap[operator_not] = operatorMongo_not
}

func isOperator(op string) bool {
	switch op {
	case operator_startBracket,
		operator_CloseBracket,
		operator_gt,
		operator_gte,
		operator_lt,
		operator_lte,
		operator_eq,
		operator_ne,
		operator_and,
		operator_or,
		operator_not_gt,
		operator_not_gte,
		operator_not_lt,
		operator_not_lte:
		return true
	default:
		return false
	}
}

// priority
func priority(op string) int {
	switch op {
	case operator_startBracket,
		operator_CloseBracket:
		return 3
	case operator_and, operator_or:
		return 2
	case operator_gt,
		operator_gte,
		operator_lt,
		operator_lte,
		operator_eq,
		operator_ne,
		operator_not_gt,
		operator_not_gte,
		operator_not_lt,
		operator_not_lte:
		return 1
	default:
		return 0
	}
}

func operatorType(op string) opType {
	switch op {
	case operator_gt,
		operator_gte,
		operator_lt,
		operator_lte,
		operator_eq,
		operator_ne:
		return opType_comparison
	case operator_and,
		operator_or:
		return opType_logical
	case operator_not_gt,
		operator_not_gte,
		operator_not_lt,
		operator_not_lte:
		return opType_NotComp
	default:
		return opType_unknown
	}
}
