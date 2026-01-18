// Package kick provides clients for the Kick API, OAuth and for processing webhooks.
//
// # Quickstart: API Client
//
//	oAuthClient, err := kick.NewOAuthClient(kickapitypes.APIClientConfig{
//		ClientID:     "your-client-id",
//		ClientSecret: "your-client-secret",
//		HTTPClient:   http.DefaultClient,
//	})
//	if err != nil {
//		log.Fatalf("could not create OAuthClient: %v", err)
//	}
//
//	apiClient, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
//		HTTPClient:   http.DefaultClient,
//	})
//	if err != nil {
//		log.Fatalf("could not create APIClient: %v", err)
//	}
//
//	accessToken, err := oAuthClient.GetAppAccessToken(context.TODO())
//	if err != nil {
//		var apiErr *kickerrors.APIError
//		if errors.As(err, &apiErr) {
//			log.Fatalf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
//		} else {
//			log.Fatalf("internal error: %v", err)
//		}
//	}
//
//	categorySearchData, err := apiClient.Category().SearchCategories(context.TODO(), accessToken, kickfilters.NewCategoriesFilter().WithNames([]string{"Software Development"}))
//
//	if err != nil {
//		var apiErr *kickerrors.APIError
//		if errors.As(err, &apiErr) {
//			log.Fatalf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
//		} else {
//			log.Fatalf("internal error: %v", err)
//		}
//	}
//
//	log.Println("Found category:", categorySearchData.Data[0].Name)
//
// # Quickstart: Webhook Client Only
//
//	webhookClient, err := kick.NewWebhookClient("your-public-key")
//	if err != nil {
//		log.Fatalf("could not create WebhookClient: %v", err)
//	}
//
//	err := webhookClient.RegisterChatMessageSentHandler(func(
//		writer http.ResponseWriter,
//		request *http.Request,
//		headers kickwebhooktypes.KickWebhookHeaders,
//		data kickwebhooktypes.ChatMessageSent,
//	) {
//		writer.WriteHeader(http.StatusOK)
//	})
//	if err != nil {
//		log.Printf("error registering chat message sent handler: %v", err)
//	}
//
//	http.HandleFunc("/", webhookClient.WebhookHandler)
//
// # Quickstart: Combined API + Webhook Client
//
//	apiClient, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
//		HTTPClient:   http.DefaultClient,
//	})
//	if err != nil {
//		log.Fatalf("could not create APIClient: %v", err)
//	}
//
//	publicKey, err := apiClient.PublicKey().GetWebhookPublicKey(context.TODO())
//	if err != nil {
//		var apiErr *kickerrors.APIError
//		if errors.As(err, &apiErr) {
//			log.Fatalf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
//		} else {
//			log.Fatalf("internal error: %v", err)
//		}
//	}
//
//	webhookClient, err := kick.NewWebhookClient(publicKey.Data.PublicKey)
//	if err != nil {
//		log.Fatalf("could not create WebhookClient: %v", err)
//	}
//
//	// Register handlers and start server as above

package kick
