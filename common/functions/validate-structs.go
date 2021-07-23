package functions

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type (
	translatorLibs struct {
		en       locales.Translator
		uni      *ut.UniversalTranslator
		trans    ut.Translator
		validate *validator.Validate
		started  bool
	}
)

var translator *translatorLibs

func ValidateStruct(s interface{}) *string {

	if translator == nil || !translator.started {
		tr := translatorLibs{}
		tr.newTranslatorLibs()
		translator = &tr
	}

	err := translator.validate.Struct(s)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			strErr := e.Translate(translator.trans)
			return &strErr
		}
	}
	return nil
}

func (tr *translatorLibs) newTranslatorLibs() {
	tr.en = en.New()
	tr.uni = ut.New(tr.en, tr.en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	tr.trans, _ = tr.uni.GetTranslator("en")

	tr.validate = validator.New()
	en_translations.RegisterDefaultTranslations(tr.validate, tr.trans)
	tr.started = true
}
