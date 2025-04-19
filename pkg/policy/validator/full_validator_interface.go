package validator

type IFullValidatorInterface interface {
	IPolicyValidator
	IStatementValidator
	IConditionValidator
}
