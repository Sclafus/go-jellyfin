package jellyfin

import (
	"encoding/json"
	"fmt"
	"io"
)

var (
	moviesIncludeFields = []string{"Path", "Name", "SortName", "ProductionYear", "Genres", "CommunityRating", "Overview", "Taglines", "MediaSources"}
)

func (c *Client) GetMovies(opts QueryOpts) ([]*Movie, error) {
	params := c.defaultParams()
	params.setIncludeTypes(mediaTypeMovie)
	params.setPaging(opts.Paging)
	params.setSorting(opts.Sort)
	params.setFilter(mediaTypeMovie, opts.Filter)
	params.enableRecursive()
	params.setIncludeFields(moviesIncludeFields...)

	resp, err := c.get(fmt.Sprintf("/Users/%s/Items", c.userID), params)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	return c.parseMovies(resp)
}

func (c *Client) parseMovies(resp io.Reader) ([]*Movie, error) {
	movies := &movies{}
	if err := json.NewDecoder(resp).Decode(&movies); err != nil {
		return nil, fmt.Errorf("decode json: %v", err)
	}

	return movies.Movies, nil
}
