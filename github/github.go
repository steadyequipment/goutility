package github

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"

	"github.com/steadyequipment/goutility"
)

func MakeClientWithOauth2StaticToken(accessToken string) *github.Client {
	token := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: accessToken,
		},
	)
	oauthClient := oauth2.NewClient(oauth2.NoContext, token)
	return github.NewClient(oauthClient)
}

func FindRepoByOrg(client *github.Client, orgName string, repoName string) (*github.Repository, error) {

	if client == nil {
		return nil, goutility.NewError("No Client provided")
	}

	options := &github.RepositoryListByOrgOptions{
	}
	var result *github.Repository = nil
	for {
		repositories, response, repoError := client.Repositories.ListByOrg(orgName, options)
		if repoError != nil {
			return nil, goutility.NewError("Error while listing repos in org '%s' in Github: %q", orgName, repoError)
		}

		for _, repo := range repositories {
			if repo.Name != nil && *(repo.Name) == repoName {
				result = repo
				break
			}
		}

		if response.NextPage == 0 {
			break
		}

		options.ListOptions.Page = response.NextPage
	}

	if result != nil {
		return result, nil
	}

	return nil, goutility.NewError("Unable to find repo '%s/%s'", orgName, repoName)
}

func getContents(client *github.Client, repository *github.Repository, path string) (*github.RepositoryContent, []*github.RepositoryContent, *github.Response, error) {
	if client == nil {
		return nil, nil, nil, goutility.NewError("No Client provided")
	}

	if repository == nil {
		return nil, nil, nil, goutility.NewError("No Repository provided")
	}

	contents, directoryContent , response, error := client.Repositories.GetContents(
		*repository.Owner.Login,
		*repository.Name,
		path,
		&github.RepositoryContentGetOptions{Ref: "heads/master"},
	)

	return contents, directoryContent, response, error
}

func GetFileContents(client *github.Client, repository *github.Repository, path string) (*github.RepositoryContent, error) {

	contents, directoryContents, _, error := getContents(client, repository, path)
	if contents == nil && directoryContents != nil {
		return nil, goutility.NewError("%s is a directory", path)
	}

	return contents, error
}

func CreateFile(client *github.Client, repository *github.Repository, path string, options *github.RepositoryContentFileOptions) (*github.RepositoryContentResponse, error)  {
	if client == nil {
		return nil, goutility.NewError("No Client provided")
	}

	if repository == nil {
		return nil, goutility.NewError("No Repository provided")
	}

	contents, _, error := client.Repositories.CreateFile(*repository.Owner.Login, *repository.Name, path, options)
	return contents, error
}

func UpdateFile(client *github.Client, repository *github.Repository, path string, options *github.RepositoryContentFileOptions) (*github.RepositoryContentResponse, error) {
	if client == nil {
		return nil, goutility.NewError("No Client provided")
	}

	if repository == nil {
		return nil, goutility.NewError("No Repository provided")
	}

	if options == nil {
		return nil, goutility.NewError("No options provided")
	}

	contents, _, error := client.Repositories.UpdateFile(*repository.Owner.Login, *repository.Name, path, options)
	return contents, error
}

func CreateOrUpdateFile(client *github.Client, repository *github.Repository, path string, options *github.RepositoryContentFileOptions) (*github.RepositoryContentResponse, error) {
	if options == nil {
		return nil, goutility.NewError("No options provided")
	}

	existingContent, _, response, getContentsError := getContents(client, repository, path)


	if response.StatusCode == 404 {
		return CreateFile(client, repository, path, options)
	} else if existingContent != nil {
		options.SHA = existingContent.SHA
		return UpdateFile(client, repository, path, options)
	} else if getContentsError != nil {
		return nil, getContentsError
	} else {
		return nil, goutility.NewError("Unable to determine status of %s to determine whether to create or update", path)
	}
}

