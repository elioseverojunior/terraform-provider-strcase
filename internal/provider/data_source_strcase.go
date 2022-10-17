package provider

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iancoleman/strcase"
)

func dataSourceStrCase() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStrCaseRead,
		Description: "Generates Strings Cases given string content.",
		Schema: map[string]*schema.Schema{
			"string": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Set `string` value.",
			},
			"delimiter": {
				Type:        schema.TypeString,
				Required:    false,
				Description: "Set `string` delimiter.",
				Optional:    true,
				Default:     "-",
			},
			"ignore": {
				Type:        schema.TypeString,
				Required:    false,
				Description: "Set `string` Ignore Case.",
				Optional:    true,
				Default:     nil,
			},
			"to_camel": {
				Type:        schema.TypeString,
				Description: "Converted `string` into Camel Case.",
				Computed:    true,
			},
			"to_delimited": {
				Type:        schema.TypeString,
				Description: "Converted `string` into Delimited Case.",
				Computed:    true,
			},
			"to_kebab": {
				Type:        schema.TypeString,
				Description: "Converted `string` into Kebab Case.",
				Computed:    true,
			},
			"to_lower_camel": {
				Type:        schema.TypeString,
				Description: "Converted `string` into Lower Camel Case.",
				Computed:    true,
			},
			"to_screaming_delimited": {
				Type:        schema.TypeString,
				Description: "Converted `string` into Screaming Delimited Case.",
				Computed:    true,
			},
			"to_screaming_kebab": {
				Type:        schema.TypeString,
				Description: "Converted `string` into Screaming Kebab Case.",
				Computed:    true,
			},
			"to_screaming_snake": {
				Type:        schema.TypeString,
				Description: "Converted `string` into Screaming Snake Case.",
				Computed:    true,
			},
			"to_snake": {
				Type:        schema.TypeString,
				Description: "Converted `string` into Snake Case.",
				Computed:    true,
			},
			"to_snake_with_ignore": {
				Type:        schema.TypeString,
				Description: "Converted `string` into Snake Case With Ignore.",
				Computed:    true,
			},
		},
	}
}

func dataSourceStrCaseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	strCaseString := d.Get("string").(string)
	strCaseIgnore := d.Get("ignore").(string)
	uint8Delimiter := uint8([]rune(d.Get("delimiter").(string))[0])

	d.SetId(getMD5Hash(strCaseString))

	if err := d.Set("to_camel", strcase.ToCamel(strCaseString)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("to_delimited", strcase.ToDelimited(strCaseString, uint8Delimiter)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("to_kebab", strcase.ToKebab(strCaseString)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("to_lower_camel", strcase.ToLowerCamel(strCaseString)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("to_screaming_delimited", strcase.ToScreamingDelimited(strCaseString, uint8Delimiter, strCaseIgnore, true)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("to_screaming_kebab", strcase.ToScreamingKebab(strCaseString)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("to_screaming_snake", strcase.ToScreamingSnake(strCaseString)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("to_snake", strcase.ToSnake(strCaseString)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("to_snake_with_ignore", strcase.ToSnakeWithIgnore(strCaseString, strCaseIgnore)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func getMD5Hash(text string) string {
	var hasher = md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
