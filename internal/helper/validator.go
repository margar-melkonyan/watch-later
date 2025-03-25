package helper

import (
	"context"
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
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
			"{field}", strings.ToLower(err.Field()),
		)
		if err.Param() != "" {
			res = strings.ReplaceAll(res, "{param}", err.Param())
		}

		validatedMessages[strings.ToLower(err.Field())] = res
	}

	return validatedMessages, nil
}
