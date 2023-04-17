package openai

import "context"

type Model struct {
	ID         string            `json:"id"`
	Object     string            `json:"object"`
	Created    int64             `json:"created"`
	OwnedBy    string            `json:"owned_by"`
	Permission []ModelPermission `json:"permission"`
	Root       string            `json:"root"`
}

type ModelPermission struct {
	ID                 string `json:"id"`
	Object             string `json:"object"`
	Created            int64  `json:"created"`
	AllowCreateEngine  bool   `json:"allow_create_engine"`
	AllowSampling      bool   `json:"allow_sampling"`
	AllowLogprobs      bool   `json:"allow_logprobs"`
	AllowSearchIndices bool   `json:"allow_search_indices"`
	AllowView          bool   `json:"allow_view"`
	AllowFineTuning    bool   `json:"allow_fine_tuning"`
	Organization       string `json:"organization"`
	IsBlocking         bool   `json:"is_blocking"`
}

// Models returns a list of models.
// Lists the currently available models, and provides basic information about each one such as the owner and availability.
func (c *Client) Models(ctx context.Context) ([]Model, error) {
	var res struct {
		Data []Model `json:"data"`
	}

	err := c.Request(ctx, "models", nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

// Model returns a model by ID.
// Retrieves a model instance, providing basic information about the model such as the owner and permissioning.
func (c *Client) Model(ctx context.Context, id string) (*Model, error) {
	res := new(Model)

	err := c.Request(ctx, "models/"+id, nil, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
