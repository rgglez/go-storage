// Copyright The go-storage Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package onedrive

// Onedrive only use OAuth2.0 authentication.

// ref:
// https://docs.microsoft.com/en-us/graph/auth/auth-concepts
// https://docs.microsoft.com/en-us/advertising/guides/authentication-oauth-get-tokens?view=bingads-13

import (
	"context"

	"github.com/goh-chunlin/go-onedrive/onedrive"
	"golang.org/x/oauth2"
)

// getClient generate new onedrive client with oauth2 token.
func getClient(ctx context.Context, token string) *onedrive.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := onedrive.NewClient(tc)

	return client
}
