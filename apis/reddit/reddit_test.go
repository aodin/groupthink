package reddit

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func decodeString(content string) *json.Decoder {
	return json.NewDecoder(bytes.NewBuffer([]byte(content)))
}

var search = decodeString(`{"kind": "Listing", "data": {"modhash": "", "children": [{"kind": "t3", "data": {"domain": "teslamotors.com", "banned_by": null, "media_embed": {}, "subreddit": "Android", "selftext_html": null, "selftext": "", "likes": null, "suggested_sort": null, "user_reports": [], "secure_media": null, "link_flair_text": null, "id": "313qxe", "gilded": 0, "archived": false, "clicked": false, "report_reasons": null, "author": "ahatzz11", "media": null, "score": 62, "approved_by": null, "over_18": false, "hidden": false, "thumbnail": "default", "subreddit_id": "t5_2qlqh", "edited": false, "link_flair_css_class": null, "author_flair_css_class": "userBlue", "downs": 0, "mod_reports": [], "secure_media_embed": {}, "saved": false, "removal_reason": null, "is_self": false, "name": "t3_313qxe", "permalink": "/r/Android/comments/313qxe/teslas_response_to_recent_smartwatches_announcing/", "stickied": false, "created": 1427922662.0, "url": "http://www.teslamotors.com/blog/announcing-tesla-model-w", "author_flair_text": "s5", "title": "Tesla's response to recent smartwatches - Announcing the Tesla Model W", "created_utc": 1427919062.0, "ups": 62, "num_comments": 6, "visited": false, "num_reports": null, "distinguished": null}}, {"kind": "t3", "data": {"domain": "teslamotors.com", "banned_by": null, "media_embed": {}, "subreddit": "teslamotors", "selftext_html": null, "selftext": "", "likes": null, "suggested_sort": null, "user_reports": [], "secure_media": null, "link_flair_text": null, "id": "313o8x", "gilded": 0, "archived": false, "clicked": false, "report_reasons": null, "author": "anontipster", "media": null, "score": 25, "approved_by": null, "over_18": false, "hidden": false, "thumbnail": "default", "subreddit_id": "t5_2s3j5", "edited": false, "link_flair_css_class": null, "author_flair_css_class": "1 f", "downs": 0, "mod_reports": [], "secure_media_embed": {}, "saved": false, "removal_reason": null, "is_self": false, "name": "t3_313o8x", "permalink": "/r/teslamotors/comments/313o8x/announcing_the_tesla_model_w/", "stickied": false, "created": 1427921706.0, "url": "http://www.teslamotors.com/blog/announcing-tesla-model-w", "author_flair_text": "Tesla Fan", "title": "Announcing the Tesla Model W", "created_utc": 1427918106.0, "ups": 25, "num_comments": 13, "visited": false, "num_reports": null, "distinguished": null}}, {"kind": "t3", "data": {"domain": "teslamotors.com", "banned_by": null, "media_embed": {}, "subreddit": "technology", "selftext_html": null, "selftext": "", "likes": null, "suggested_sort": null, "user_reports": [], "secure_media": null, "link_flair_text": "Transport", "id": "313qim", "gilded": 0, "archived": false, "clicked": false, "report_reasons": null, "author": "ahatzz11", "media": null, "score": 0, "approved_by": null, "over_18": false, "hidden": false, "thumbnail": "", "subreddit_id": "t5_2qh16", "edited": false, "link_flair_css_class": "general", "author_flair_css_class": null, "downs": 0, "mod_reports": [], "secure_media_embed": {}, "saved": false, "removal_reason": null, "is_self": false, "name": "t3_313qim", "permalink": "/r/technology/comments/313qim/announcing_the_tesla_model_w/", "stickied": false, "created": 1427922529.0, "url": "http://www.teslamotors.com/blog/announcing-tesla-model-w", "author_flair_text": null, "title": "Announcing the Tesla Model W", "created_utc": 1427918929.0, "ups": 0, "num_comments": 3, "visited": false, "num_reports": null, "distinguished": null}}, {"kind": "t3", "data": {"domain": "teslamotors.com", "banned_by": null, "media_embed": {}, "subreddit": "news", "selftext_html": null, "selftext": "", "likes": null, "suggested_sort": null, "user_reports": [], "secure_media": null, "link_flair_text": null, "id": "316uxy", "gilded": 0, "archived": false, "clicked": false, "report_reasons": null, "author": "turtlewong", "media": null, "score": 2, "approved_by": null, "over_18": false, "hidden": false, "thumbnail": "", "subreddit_id": "t5_2qh3l", "edited": false, "link_flair_css_class": null, "author_flair_css_class": null, "downs": 0, "mod_reports": [], "secure_media_embed": {}, "saved": false, "removal_reason": null, "is_self": false, "name": "t3_316uxy", "permalink": "/r/news/comments/316uxy/tesla_trolling_apple/", "stickied": false, "created": 1427985549.0, "url": "http://www.teslamotors.com/blog/announcing-tesla-model-w", "author_flair_text": null, "title": "Tesla trolling Apple", "created_utc": 1427981949.0, "ups": 2, "num_comments": 0, "visited": false, "num_reports": null, "distinguished": null}}, {"kind": "t3", "data": {"domain": "teslamotors.com", "banned_by": null, "media_embed": {}, "subreddit": "aprilfools", "selftext_html": null, "selftext": "", "likes": null, "suggested_sort": null, "user_reports": [], "secure_media": null, "link_flair_text": null, "id": "313pnh", "gilded": 0, "archived": false, "clicked": false, "report_reasons": null, "author": "CashOverAss", "media": null, "score": 1, "approved_by": null, "over_18": false, "hidden": false, "thumbnail": "", "subreddit_id": "t5_2qhpu", "edited": false, "link_flair_css_class": null, "author_flair_css_class": null, "downs": 0, "mod_reports": [], "secure_media_embed": {}, "saved": false, "removal_reason": null, "is_self": false, "name": "t3_313pnh", "permalink": "/r/aprilfools/comments/313pnh/announcing_the_tesla_model_w/", "stickied": false, "created": 1427922200.0, "url": "http://www.teslamotors.com/blog/announcing-tesla-model-w", "author_flair_text": null, "title": "Announcing the Tesla Model W", "created_utc": 1427918600.0, "ups": 1, "num_comments": 0, "visited": false, "num_reports": null, "distinguished": null}}, {"kind": "t3", "data": {"domain": "teslamotors.com", "banned_by": null, "media_embed": {}, "subreddit": "funny", "selftext_html": null, "selftext": "", "likes": null, "suggested_sort": null, "user_reports": [], "secure_media": null, "link_flair_text": null, "id": "318ebo", "gilded": 0, "archived": false, "clicked": false, "report_reasons": null, "author": "BufloSolja", "media": null, "score": 1, "approved_by": null, "over_18": false, "hidden": false, "thumbnail": "default", "subreddit_id": "t5_2qh33", "edited": false, "link_flair_css_class": null, "author_flair_css_class": null, "downs": 0, "mod_reports": [], "secure_media_embed": {}, "saved": false, "removal_reason": null, "is_self": false, "name": "t3_318ebo", "permalink": "/r/funny/comments/318ebo/announcing_the_new_tesla_model_w/", "stickied": false, "created": 1428010627.0, "url": "http://www.teslamotors.com/blog/announcing-tesla-model-w", "author_flair_text": null, "title": "Announcing the New Tesla Model W", "created_utc": 1428007027.0, "ups": 1, "num_comments": 0, "visited": false, "num_reports": null, "distinguished": null}}], "after": null, "before": null}}`)

var comments = decodeString(`[{"kind": "Listing", "data": {"modhash": "", "children": [{"kind": "t3", "data": {"domain": "teslamotors.com", "banned_by": null, "media_embed": {}, "subreddit": "Android", "selftext_html": null, "selftext": "", "likes": null, "suggested_sort": null, "user_reports": [], "secure_media": null, "link_flair_text": null, "id": "313qxe", "gilded": 0, "archived": false, "clicked": false, "report_reasons": null, "author": "ahatzz11", "media": null, "score": 64, "approved_by": null, "over_18": false, "hidden": false, "num_comments": 6, "thumbnail": "default", "subreddit_id": "t5_2qlqh", "edited": false, "link_flair_css_class": null, "author_flair_css_class": "userBlue", "downs": 0, "secure_media_embed": {}, "saved": false, "removal_reason": null, "stickied": false, "is_self": false, "permalink": "/r/Android/comments/313qxe/teslas_response_to_recent_smartwatches_announcing/", "name": "t3_313qxe", "created": 1427922662.0, "url": "http://www.teslamotors.com/blog/announcing-tesla-model-w", "author_flair_text": "s5", "title": "Tesla's response to recent smartwatches - Announcing the Tesla Model W", "created_utc": 1427919062.0, "distinguished": null, "upvote_ratio": 0.8, "mod_reports": [], "visited": false, "num_reports": null, "ups": 64}}], "after": null, "before": null}}, {"kind": "Listing", "data": {"modhash": "", "children": [{"kind": "t1", "data": {"subreddit_id": "t5_2qlqh", "banned_by": null, "removal_reason": null, "link_id": "t3_313qxe", "likes": null, "replies": "", "user_reports": [], "saved": false, "id": "cpy6x8q", "gilded": 0, "archived": false, "report_reasons": null, "author": "SirFadakar", "parent_id": "t3_313qxe", "score": 15, "approved_by": null, "controversiality": 0, "body": "This one was actually hilarious because it's takes obvious to the next level. The anti-prank.", "edited": false, "author_flair_css_class": null, "downs": 0, "body_html": "&lt;div class=\"md\"&gt;&lt;p&gt;This one was actually hilarious because it&amp;#39;s takes obvious to the next level. The anti-prank.&lt;/p&gt;\n&lt;/div&gt;", "subreddit": "Android", "score_hidden": false, "name": "t1_cpy6x8q", "created": 1427924835.0, "author_flair_text": null, "created_utc": 1427921235.0, "distinguished": null, "mod_reports": [], "num_reports": null, "ups": 15}}, {"kind": "more", "data": {"count": 5, "parent_id": "t3_313qxe", "children": ["cpycofx", "cpy5r4q", "cpynhus", "cpya226", "cpymjn2"], "id": "cpycofx", "name": "t1_cpycofx"}}], "after": null, "before": null}}]`)

func TestReddit_unmarshalSearch(t *testing.T) {
	// Test the unmarshal of the example query
	results, err := unmarshalSearch(search)
	require.Nil(t, err)
	require.Equal(t, 6, len(results))

	// The first result
	assert.Equal(t, "313qxe", results[0].Data.ID)
	assert.Equal(t,
		"Tesla's response to recent smartwatches - Announcing the Tesla Model W",
		results[0].Data.Title,
	)
	assert.Equal(t, 6, results[0].Data.NumberOfComments)
	assert.Equal(t, 62, results[0].Data.Score)
	assert.Equal(t, "teslamotors.com", results[0].Data.Domain)
	assert.Equal(t,
		time.Date(2015, 4, 1, 20, 11, 2, 0, time.UTC),
		results[0].Data.Created.AsTime(),
	)
}

func TestReddit_unmarshalComments(t *testing.T) {
	// Test the unmarshal of the example query
	results, err := unmarshalComments(comments)
	require.Nil(t, err)
	require.Equal(t, 2, len(results))

	assert.Equal(t,
		"This one was actually hilarious because it's takes obvious to the next level. The anti-prank.",
		results[0].Data.Body,
	)
	assert.Equal(t,
		time.Date(2015, 4, 1, 20, 47, 15, 0, time.UTC),
		results[0].Data.Created.AsTime(),
	)
}
