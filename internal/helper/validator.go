package helper

import (
	"context"
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	"github.com/margar-melkonyan/watch-later.git/internal/lang/eng"
	"github.com/margar-melkonyan/watch-later.git/internal/lang/ru"
)

func getValidationMessages(locale string) map[string]string {
	switch locale {
	case "ru":
		return ru.GetMessages()
	default:
		return eng.GetMessages()
	}
}

func getAttribute(locale string, attribute string) string {
	switch locale {
	case "ru":
		return ru.GetAttribute(attribute)
	default:
		return eng.GetAttribute(attribute)
	}
}

func LocalizedValidationMessages(
	ctx context.Context,
	errs validator.ValidationErrors,
) (map[string]string, error) {
	locale, ok := ctx.Value("locale").(string)
	if !ok {
		return nil, errors.New("value is not correct")
	}
	if locale == "" {
		return nil, errors.New("locale is not set")
	}
	validationMessages := getValidationMessages(locale)
	validatedMessages := make(map[string]string)

	for _, err := range errs {
		var res string
		res = strings.ReplaceAll(
			validationMessages[err.Tag()],
			"{field}", getAttribute(locale, strcase.ToSnake(err.Field())),
		)
		if err.Param() != "" {
			res = strings.ReplaceAll(
				res,
				"{param}", getAttribute(locale, strcase.ToSnake(err.Param())),
			)
		}

		validatedMessages[strcase.ToSnake(err.Field())] = res
	}

	return validatedMessages, nil
}
