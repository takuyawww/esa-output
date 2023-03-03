# esa-output

## usage
* esaのAPIを使用して、CSV出力するためのツールです。
* [esaのAPIドキュメント](https://docs.esa.io/posts/102)に`esa accept up to 75 requests per user per 15 minutes`とあるように、APIにはリクエスト制限があります。そちらの制限にかかる場合の動作確認はしていません。

```sh
$ ./esa-output --help
Flags:
  --[no-]help                  Show context-sensitive help (also try --help-long and --help-man).
  --team=TEAM                  your team name (*required*)
  --access-token=ACCESS-TOKEN  your esa personal access token (*required*)
  --per-page=100               per page default 100 (in 20 ~ 100)
  --sort-posts="number"        sort by posts default number (sort by updated, created, number, stars,
                               watches, comments, best_match)
  --sort-members="number"      sort by members default joined (sort by posts_count, joined, last_accessed)
  --order="asc"                order by default asc (order by desc or asc)

# output esa.csv at current directory
$ ./esa-output --team=your-team-name --access-token=your-access-token
```