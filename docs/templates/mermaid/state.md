```mermaid
stateDiagram-v2
    state "未着手" as todo
    state "対応中" as in_progress
    state "レビュー中" as in_review
    state "レビュー修正" as in_fix
    state "承認済" as approve
    state "クローズ" as close

    [*] --> todo : 起票
    todo --> in_progress : 着手
    in_progress --> in_review : レビュー依頼
    in_review --> in_fix : 修正依頼
    in_fix --> in_review : レビュー依頼
    in_review --> approve : 承認
    approve --> close : 対応完了
    close --> [*]
```
