// Package kick provides clients for the Kick API and for processing webhooks.
//
// # Quickstart: API Client Only
//
//	apiClient, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
//		ClientID:     "your-client-id",
//		ClientSecret: "your-client-secret",
//		HTTPClient:   http.DefaultClient,
//	})
//	if err != nil {
//		log.Fatalf("could not create APIClient: %v", err)
//	}
//
//	currentUser, err := apiClient.User.GetCurrentUser(context.TODO(), "user-access-token")
//	if err != nil {
//		log.Printf("could not get current user: %v", err)
//	}
//	log.Println("Logged in as:", currentUser.Data[0].Username)
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
//		apiClient, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
//			ClientID:     "your-client-id",
//			ClientSecret: "your-client-secret",
//			HTTPClient:   http.DefaultClient,
//		})
//		if err != nil {
//			log.Fatalf("could not create APIClient: %v", err)
//		}
//
//		publicKey, err := apiClient.PublicKey.GetWebhookPublicKey(context.TODO())
//		if err != nil {
//			log.Fatalf("could not get the public key for the webhook: %v", err)
//		}
//
//		webhookClient, err := kick.NewWebhookClient(publicKey.Data.PublicKey)
//		if err != nil {
//			log.Fatalf("could not create WebhookClient: %v", err)
//		}
//
//		// Register handlers and start server as above
//	 //
//	 // Note: Provide a dummy ClientID and ClientSecret if you just want to automate
//		// creation of the webhook client with the public key without using any other part
//		// of the API Client. Credentials are not required for that route
package kick
