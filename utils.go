package isvosa

import "sync"

var (
    baseURL string = "https://api.telegram.org"
    previousUpdate Update 
    wg sync.WaitGroup
    previousID int = 0
)

type Bot struct {
    Token string
}

type Updates struct {
    Update []Update `json:"result"`
}

type Update struct {
    ID int `json:"update_id"`
    Message Message `json:"message,omitempty"`
    EditedMessage Message `json:"edited_message,omitempty"`
    ChannelPost Message `json:"channel_post,omitempty"`
    EditedChannelPost Message `json:"edited_channel_post,omitempty"`
    InlineQuery InlineQuery `json:"inline_query,omitempty"`
    ChosenInlineResult ChosenInlineResult `json:"chosen_inline_result,omitempty"`
    CallbackQuery CallbackQuery `json:"callback_query,omitempty"`
    ShippingQuery ShippingQuery `json:"shipping_query,omitempty"`
    PreCheckoutQuery PreCheckoutQuery `json:"pre_checkout_query,omitempty"`
    Poll Poll `json:"poll,omitempty"`
    PollAnswer PollAnswer `json:"poll_answer,omitempty"`
    MyChatMember ChatMemberUpdated `json:"my_chat_member"`
    ChatMember ChatMemberUpdated `json:"chat_member"`
    ChatJoinRequest ChatJoinRequest `json:"chat_join_request"`
    Command string
    Args []string
}

type Message struct {
    ID int `json:"message_id"`
    From User `json:"from,omitempty"`
    SenderChat Chat `json:"sender_chat,omitempty"`
    Date int `json:"date"`
    Chat Chat `json:"chat"`
    ForwardFrom User `json:"forward_from,omitempty"`
    ForwardFromChat Chat `json:"forward_from_chat,omitempty"`
    ForwardFromMessageID int `json:"forward_from_message_id,omitempty"`
    ForwardSignature string `json:"forward_signature,omitempty"`
    ForwardSenderName string `json:"forward_sender_name,omitempty"`
    ForwardDate int `json:"forward_date,omitempty"`
    IsAutomaticForward bool `json:"is_automatic_forward,omitempty"`
    ReplyToMessage Message `json:"reply_to_message,omitempty"`
    ViaBot User `json:"via_bot,omitempty"`
    EditDate int `json:"edit_date,omitempty"`
    HasProtectedContent bool `json:"has_protected_content"`
    MediaGroupID string `json:"media_group_id,omitempty"`
    AuthorSignature string `json:"author_signature,omitempty"`
    Text string `json:"text,omitempty"`
    Entities []MessageEntity `json:"entities,omitempty"`
    Animation Animation `json:"animation,omitempty"`
    Audio Audio `json:"audio,omitempty"`
    Document Document `json:"document,omitempty"`
    Photo []PhotoSize `json:"photo,omitempty"`
    Sticker Sticker `json:"sticker,omitempty"`
    Video Video `json:"video,omitempty"`
    VideoNote VideoNote `json:"video_note,omitempty"`
    Voice Voice `json:"voice,omitempty"`
    Caption string `json:"caption,omitempty"`
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    Contact Contact `json:"contact,omitempty"`
    Dice Dice `json:"dice,omitempty"`
    Game Game `json:"game,omitempty"`
    Poll Poll `json:"poll,omitempty"`
    Venue Venue `json:"venue,omitempty"`
    Location Location `json:"location,omitempty"`
    NewChatMembers []User `json:"new_chat_members,omitempty"`
    LeftChatMember User `json:"left_chat_member,omitempty"`
    NewChatTitle string `json:"new_chat_title,omitempty"`
    NewChatPhoto []PhotoSize `json:"new_chat_photo,omitempty"`
    DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`
    GroupChatCreated bool `json:"group_chat_created,omitempty"`
    SupergroupChatCreated bool `json:"supergroup_chat_created,omitempty"`
    ChannelChatCreated bool `json:"channel_chat_created,omitempty"`
    MessageAutoDeleteTimerChanged []MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
    MigrateToChatID int `json:"migrate_to_chat_id,omitempty"`
    MigrateFromChatID int `json:"migrate_from_chat_id,omitempty"`
    PinnedMessage Message `json:"pinned_message"`
    Invoice Invoice `json:"invoice,omitempty"`
    SuccessfulPayment SuccessfulPayment `json:"successful_payment,omitempty"`
    ConnectedWebsite string `json:"connected_website,omitempty"`
    PassportData PassportDate `json:"password_data,omitempty"`
    ProximityAlertTriggered ProximityAlertTriggered `json:"proximity_alert_triggered"`
    VoiceChatScheduled VoiceChatScheduled `json:"voice_chat_scheduled,omitempty"`
    VoiceChatStarted VoiceChatStarted `json:"voice_chat_started,omitempty"`
    VoiceChatEnded VoiceChatEnded `json:"voice_chat_ended,omitempty"`
    VoiceChatParticipantsInvited VoiceChatParticipantsInvited `json:"voice_chat_participants_invited,omitempty"`
    ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

type Chat struct {
    ID int `json:"id"`
    Type string `json:"type"`
    Title string `json:"title,omitempty"`
    Username string `json:"username,omitempty"`
    FirstName string `json:"first_name,omitempty"`
    LastName string `json:"last_name,omitempty"`
    Photo ChatPhoto `json:"photo,omitempty"`
    Bio string `json:"bio,omitempty"`
    HasPrivateForwards bool `json:"has_private_forwards,omitempty"`
    Description string `json:"description,omitempty"`
    InviteLink string `json:"invite_link,omitempty"`
    PinnedMessage Message `json:"pinned_message,omitempty"`
    Permissions ChatPermissions `json:"permissions,omitempty"`
    SlowModeDelay int `json:"slow_mode_delay,omitempty"`
    MessageAutoDeleteTime int `json:"message_auto_delete_time,omitempty"`
    HasProtectedContent bool `json:"has_protected_content,omitempty"`
    SticketSetName string `json:"sticker_set_name,omitempty"`
    CanSetStickerSet bool `json:"can_set_sticker_set,omitempty"`
    LinkedChatID int `json:"linked_chat_id,omitempty"`
    Location ChatLocation `json:"location,omitempty"`
}

type InlineQuery struct {
    ID string `json:"id"`
    From User `json:"from"`
    Query string `json:"query"`
    Offset string `json:"offset"`
    ChatType string `json:"chat_type,omitempty"`
    Location Location `json:"location,omitempty"`
}

type User struct {
    ID int `json:"id"`
    IsBot bool `json:"is_bot"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name,omitempty"`
    Username string `json:"username,omitempty"`
    LanguageCode string `json:"language_code,omitempty"`
    CanJoinGroups bool `json:"can_join_groups,omitempty"`
    CanReadAllGroupMessages bool `json:"can_read_all_group_messages,omitempty"`
    SupportInlineQueries bool `json:"support_inline_queries,omitempty"`
}

type Location struct {
    Longitude float64 `json:"longitude"`
    Latitude float64 `json:"latitude"`
    HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
    LivePeriod int `json:"live_period,omitempty"`
    Heading int `json:"heading,omitempty"`
    ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
}

type ChosenInlineResult struct {
    ResultID string `json:"result_id"`
    From User `json:"from"`
    Location Location `json:"location,omitempty"`
    InlineMessageID string `json:"inline_message_id,omitempty"`
    Query string `json:"query"`
}

type CallbackQuery struct {
    ID string `json:"id"`
    From User `json:"from"`
    Message Message `json:"Message"`
    InlineMessageID string `json:"inline_message_id,omitempty"`
    ChatInstance string `json:"chat_instance"`
    Data string `json:"data,omitempty"`
    GameShortName string `json:"game_short_name,omitempty"`
}

type ShippingQuery struct {
    ID string `json:"id"`
    From User `json:"from"`
    InvoicePayload string `json:"invoice_payload"`
    ShippingAddress ShippingAddress `json:"shipping_address"`
}

type ShippingAddress struct {
    CountryCode string `json:"country_code"`
    State string `json:"state"`
    City string `json:"city"`
    StreetLine1 string `json:"street_line1"`
    StreetLine2 string `json:"street_line2"`
    PostCode string `json:"post_code"`
}

type PreCheckoutQuery struct {
    ID string `json:"id"`
    From User `json:"from"`
    Currency string `json:"currency"`
    TotalAmount int `json:"total_amount"`
    InvoicePayload string `json:"invoice_payload"`
    ShippingOptionID string `json:"shipping_option_id,omitempty"`
    OrderInfo string `json:"order_info,omitempty"`
}

type Poll struct {
    ID string `json:"id"`
    Question string `json:"question"`
    Options []PollOption `json:"options"`
    TotalVoterCount int `json:"total_voter_count"`
    IsClosed bool `json:"is_closed"`
    IsAnonymous bool `json:"is_anonymous"`
    Type string `json:"type"`
    AllowsMultipleAnswers bool `json:"allows_multiple_answers"`
    CorrectOptionID int `json:"correct_option_id,omitempty"`
    Explanation string `json:"explanation,omitempty"`
    ExplanationEntities []MessageEntity `json:"explanation_entities,omitempty"`
    OpenPeriod int `json:"open_period,omitempty"`
    CloseDate int `json:"close_date,omitempty"`
}

type PollOption struct {
    Text string `json:"text"`
    VoterCount int `json:"voter_count"`
}

type MessageEntity struct {
    Type string `json:"type"`
    Offset int `json:"offset"`
    Length int `json:"length"`
    URL string `json:"url,omitempty"`
    User User `json:"user,omitempty"`
    Language string `json:"language,omitempty"`
}

type PollAnswer struct {
    PollID string `json:"poll_id"`
    User User `json:"user"`
    OptionIDs []int `json:"option_ids"`
}

type ChatMemberUpdated struct {
    Chat Chat `json:"chat"`
    From User `json:"from"`
    Date int `json:"date"`
    OldChatMember interface{} `json:"old_chat_member"`
    NewChatMember interface{} `json:"new_chat_member"`
    InviteLink ChatInviteLink `json:"invite_link,omitempty"`
}

type ChatMemberOwner struct {
    Status string `json:"status"`
    User User `json:"user"`
    IsAnonymous bool `json:"is_anonymous"`
    CustomTitle string `json:"custom_title,omitempty"`
}

type ChatMemberAdministrator struct {
    Status string `json:"status"`
    User User `json:"user"`
    CanBeEdited bool `json:"can_be_edited"`
    IsAnonymous bool `json:"is_anonymous"`
    CanManageChat bool `json:"can_manage_chat"`
    CanDeleteMessages bool `json:"can_delete_messages"`
    CanManageVoiceChats bool `json:"can_manage_voice_chats"`
    CanRestrictMembers bool `json:"can_restrict_members"`
    CanPromoteMembers bool `json:"can_promote_members"`
    CanChangeInfo bool `json:"can_change_info"`
    CanInviteUsers bool `json:"can_invite_users"`
    CanPostMessages bool `json:"can_post_messages,omitempty"`
    CanEditMessages bool `json:"can_edit_messages,omitempty"`
    CanPinMessages bool `json:"can_pin_messages,omitempty"`
    CustomTitle string `json:"custom_title,omitempty"`
}

type ChatMemberMember struct {
    Status string `json:"status"`
    User User `json:"user"`
}

type ChatMemberRestricted struct {
    Status string `json:"status"`
    User User `json:"user"`
    IsMember bool `json:"is_member"`
    CanChangeInfo bool `json:"can_change_info"`
    CanInviteUsers bool `json:"can_invite_users"`
    CanPinMessages bool `json:"can_pin_messages"`
    CanSendMessages bool `json:"can_send_messages"`
    CanSendMediaMessages bool `json:"can_send_media_messages"`
    CanSendPolls bool `json:"can_send_polls"`
    CanSendOtherMessages bool `json:"can_send_other_messages"`
    CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
    UntilDate int `json:"until_date"`
}

type ChatMemberLeft struct {
    Status string `json:"status"`
    User User `json:"user"`
}

type ChatMemberBanned struct {
    Status string `json:"status"`
    User User `json:"user"`
    UntilDate int `json:"until_date"`
}

type ChatJoinRequest struct {
    Chat Chat `json:"chat"`
    From User `json:"from"`
    Date int `json:"date"`
    Bio string `json:"bio,omitempty"`
    InviteLink ChatInviteLink `json:"invite_link,omitempty"`
}

type ChatInviteLink struct {
    InviteLink string `json:"invite_link"`
    Creator User `json:"creator"`
    CreatesJoinRequest bool `json:"creates_join_request"`
    IsPrimary bool `json:"is_primary"`
    IsRevoked bool `json:"is_revoked"`
    Name string `json:"name,omitempty"`
    ExpireDate int `json:"expire_date,omitempty"`
    MemberLimit int `json:"member_limit,omitempty"`
    PendingJoinRequestCount int `json:"pending_join_request_count,omitempty"`
}

type Me struct {
    Me struct {
        ID int `json:"id"`
        IsBot bool `json:"is_bot"`
        FirstName string `json:"first_name"`
        Username string `json:"username"`
        CanJoinGroups bool `json:"can_join_groups"`
        CanReadAllGroupMessages bool `json:"can_read_all_group_messages"`
        SupportInlineQueries bool `json:"support_inline_queries"`
    } `json:"result"`
}

type Msg struct {
    ChatID int `json:"chat_id"`
    Text string `json:"text"`
    ParseMode string `json:"parse_mode,omitempty"`
    Entities []MessageEntity `json:"entities,omitempty"`
    DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyToMessageID int `json:"reply_to_message_id,omitempty"`
    AllowSendingWithoutReply bool `json:"allow_sending_without_reply"`
    ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

type Photo struct {
    ChatID int `json:"chat_id"`
    Photo interface{} `json:"photo"`
    Caption string `json:"caption,omitempty"`
    ParseMode string `json:"parse_mode,omitempty"`
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyToMessageID int `json:"reply_to_message_id,omitempty"`
    AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
    ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

type Audio struct {
    ChatID int `json:"chat_id"`
    Audio interface{} `json:"audio"`
    Caption string `json:"caption,omitempty"`
    ParseMode string `json:"parse_mode,omitempty"`
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    Duration int `json:"duration,omitempty"`
    Performer string `json:"performer,omitempty"`
    Title string `json:"title,omitempty"`
    Thumb interface{} `json:"thumb,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyToMessageID int `json:"reply_to_message_id,omitempty"`
    AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
    ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

type Document struct {
    ChatID int `json:"chat_id"`
    Document interface{} `json:"document"`
    Thumb interface{} `json:"thumb,omitempty"`
    Caption string `json:"caption,omitempty"`
    ParseMode string `json:"parse_mode,omitempty"`
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    DisableContentTypeDetection bool `json:"disable_content_tyoe_detection,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyToMessageID int `json:"reply_to_message_id,omitempty"`
    AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
    ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

type Video struct {
    ChatID int `json:"chat_id"`
    Video interface{} `json:"video"`
    Duration int `json:"duration,omitempty"`
    Width int `json:"width,omitempty"`
    Height int `json:"height,omitempty"`
    Thumb interface{} `json:"thumb,omitempty"`
    Caption string `json:"caption,omitempty"`
    ParseMode string `json:"parse_mode,omitempty"`
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    SupportsStreaming bool `json:"supports_streaming,omitempty"`
    DisableContentTypeDetection bool `json:"disable_content_tyoe_detection,omitempty"`
    DisableNotification bool `json:"disable_notification,omitempty"`
    ProtectContent bool `json:"protect_content,omitempty"`
    ReplyToMessageID int `json:"reply_to_message_id,omitempty"`
    AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
    ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}
