package core

type Text string

const (
	// Custom Server
	TextCustomServer Text = "#CUSTOM_SERVER"
	// Collection
	TextCollection Text = "#COLLECTION"
	// Server
	TextServer Text = "#SERVER"
	// Address
	TextAddress Text = "#ADDRESS"
	// Expiration
	TextExpiration Text = "#EXPIRATION"
	// Storage
	TextStorage Text = "#STORAGE"
	// Log
	TextLog Text = "#LOG"
	// Inquiry
	TextInquiry Text = "#INQUIRY"
	// Banner
	TextBanner Text = "#BANNER"
	// Post
	TextPost Text = "#POST"
	// Playlist
	TextPlaylist Text = "#PLAYLIST"
	// Level
	TextLevel Text = "#LEVEL"
	// Skin
	TextSkin Text = "#SKIN"
	// Background
	TextBackground Text = "#BACKGROUND"
	// SFX
	TextEffect Text = "#EFFECT"
	// Particle
	TextParticle Text = "#PARTICLE"
	// Engine
	TextEngine Text = "#ENGINE"
	// Replay
	TextReplay Text = "#REPLAY"
	// User
	TextUser Text = "#USER"
	// Room
	TextRoom Text = "#ROOM"
	// Thumbnail
	TextPostThumbnail Text = "#POST_THUMBNAIL"
	// Thumbnail
	TextPlaylistThumbnail Text = "#PLAYLIST_THUMBNAIL"
	// Cover
	TextLevelCover Text = "#LEVEL_COVER"
	// BGM
	TextLevelBgm Text = "#LEVEL_BGM"
	// Preview
	TextLevelPreview Text = "#LEVEL_PREVIEW"
	// Data
	TextLevelData Text = "#LEVEL_DATA"
	// Thumbnail
	TextSkinThumbnail Text = "#SKIN_THUMBNAIL"
	// Data
	TextSkinData Text = "#SKIN_DATA"
	// Texture
	TextSkinTexture Text = "#SKIN_TEXTURE"
	// Thumbnail
	TextBackgroundThumbnail Text = "#BACKGROUND_THUMBNAIL"
	// Image
	TextBackgroundImage Text = "#BACKGROUND_IMAGE"
	// Data
	TextBackgroundData Text = "#BACKGROUND_DATA"
	// Configuration
	TextBackgroundConfiguration Text = "#BACKGROUND_CONFIGURATION"
	// Thumbnail
	TextEffectThumbnail Text = "#EFFECT_THUMBNAIL"
	// Data
	TextEffectData Text = "#EFFECT_DATA"
	// Audio
	TextEffectAudio Text = "#EFFECT_AUDIO"
	// Thumbnail
	TextParticleThumbnail Text = "#PARTICLE_THUMBNAIL"
	// Data
	TextParticleData Text = "#PARTICLE_DATA"
	// Texture
	TextParticleTexture Text = "#PARTICLE_TEXTURE"
	// Thumbnail
	TextEngineThumbnail Text = "#ENGINE_THUMBNAIL"
	// Play Data
	TextEnginePlaydata Text = "#ENGINE_PLAYDATA"
	// Watch Data
	TextEngineWatchdata Text = "#ENGINE_WATCHDATA"
	// Preview Data
	TextEnginePreviewdata Text = "#ENGINE_PREVIEWDATA"
	// Tutorial Data
	TextEngineTutorialdata Text = "#ENGINE_TUTORIALDATA"
	// ROM
	TextEngineRom Text = "#ENGINE_ROM"
	// Configuration
	TextEngineConfiguration Text = "#ENGINE_CONFIGURATION"
	// Data
	TextReplayData Text = "#REPLAY_DATA"
	// Configuration
	TextReplayConfiguration Text = "#REPLAY_CONFIGURATION"
	// Cover
	TextRoomCover Text = "#ROOM_COVER"
	// BGM
	TextRoomBgm Text = "#ROOM_BGM"
	// Preview
	TextRoomPreview Text = "#ROOM_PREVIEW"
	// Grade
	TextGrade Text = "#GRADE"
	// Arcade Score
	TextArcadeScore Text = "#ARCADE_SCORE"
	// Accuracy Score
	TextAccuracyScore Text = "#ACCURACY_SCORE"
	// Combo
	TextCombo Text = "#COMBO"
	// Perfect
	TextPerfect Text = "#PERFECT"
	// Great
	TextGreat Text = "#GREAT"
	// Good
	TextGood Text = "#GOOD"
	// Miss
	TextMiss Text = "#MISS"
	// Judgment
	TextJudgment Text = "#JUDGMENT"
	// Accuracy
	TextAccuracy Text = "#ACCURACY"
	// Filter
	TextFilter Text = "#FILTER"
	// Sort
	TextSort Text = "#SORT"
	// Keywords
	TextKeywords Text = "#KEYWORDS"
	// Name
	TextName Text = "#NAME"
	// Source
	TextSource Text = "#SOURCE"
	// Rating
	TextRating Text = "#RATING"
	// Minimum Rating
	TextRatingMinimum Text = "#RATING_MINIMUM"
	// Maximum Rating
	TextRatingMaximum Text = "#RATING_MAXIMUM"
	// Title
	TextTitle Text = "#TITLE"
	// Subtitle
	TextSubtitle Text = "#SUBTITLE"
	// Artists
	TextArtists Text = "#ARTISTS"
	// Time
	TextTime Text = "#TIME"
	// Author
	TextAuthor Text = "#AUTHOR"
	// Coauthor
	TextCoauthor Text = "#COAUTHOR"
	// Collaborator
	TextCollaborator Text = "#COLLABORATOR"
	// Description
	TextDescription Text = "#DESCRIPTION"
	// Genre
	TextGenre Text = "#GENRE"
	// Type
	TextType Text = "#TYPE"
	// Category
	TextCategory Text = "#CATEGORY"
	// Status
	TextStatus Text = "#STATUS"
	// Language
	TextLanguage Text = "#LANGUAGE"
	// Difficulty
	TextDifficulty Text = "#DIFFICULTY"
	// Version
	TextVersion Text = "#VERSION"
	// Length
	TextLength Text = "#LENGTH"
	// Minimum Length
	TextLengthMinimum Text = "#LENGTH_MINIMUM"
	// Maximum Length
	TextLengthMaximum Text = "#LENGTH_MAXIMUM"
	// Additional Information
	TextAdditionalInformation Text = "#ADDITIONAL_INFORMATION"
	// Timezone
	TextTimezone Text = "#TIMEZONE"
	// Region
	TextRegion Text = "#REGION"
	// Tag
	TextTag Text = "#TAG"
	// Include Tag
	TextIncludeTag Text = "#INCLUDE_TAG"
	// Exclude Tag
	TextExcludeTag Text = "#EXCLUDE_TAG"
	// Content
	TextContent Text = "#CONTENT"
	// Comment
	TextComment Text = "#COMMENT"
	// Message
	TextMessage Text = "#MESSAGE"
	// Notification
	TextNotification Text = "#NOTIFICATION"
	// Role
	TextRole Text = "#ROLE"
	// Permission
	TextPermission Text = "#PERMISSION"
	// Level Speed
	TextSpeed Text = "#SPEED"
	// Mirror Level
	TextMirror Text = "#MIRROR"
	// Random
	TextRandom Text = "#RANDOM"
	// Hidden
	TextHidden Text = "#HIDDEN"
	// Strict Judgment
	TextJudgmentStrict Text = "#JUDGMENT_STRICT"
	// Loose Judgment
	TextJudgmentLoose Text = "#JUDGMENT_LOOSE"
	// Auto SFX
	TextEffectAuto Text = "#EFFECT_AUTO"
	// Haptic
	TextHaptic Text = "#HAPTIC"
	// Stage
	TextStage Text = "#STAGE"
	// Stage Position
	TextStagePosition Text = "#STAGE_POSITION"
	// Stage Size
	TextStageSize Text = "#STAGE_SIZE"
	// Stage Rotation
	TextStageRotation Text = "#STAGE_ROTATION"
	// Stage Direction
	TextStageDirection Text = "#STAGE_DIRECTION"
	// Stage Transparency
	TextStageAlpha Text = "#STAGE_ALPHA"
	// Stage Animation
	TextStageAnimation Text = "#STAGE_ANIMATION"
	// Stage Tilt
	TextStageTilt Text = "#STAGE_TILT"
	// Vertical Stage Cover
	TextStageCoverVertical Text = "#STAGE_COVER_VERTICAL"
	// Horizontal Stage Cover
	TextStageCoverHorizontal Text = "#STAGE_COVER_HORIZONTAL"
	// Stage Cover Transparency
	TextStageCoverAlpha Text = "#STAGE_COVER_ALPHA"
	// Lock Stage Aspect Ratio
	TextStageAspectratioLock Text = "#STAGE_ASPECTRATIO_LOCK"
	// Stage Effect
	TextStageEffect Text = "#STAGE_EFFECT"
	// Stage Effect Position
	TextStageEffectPosition Text = "#STAGE_EFFECT_POSITION"
	// Stage Effect Size
	TextStageEffectSize Text = "#STAGE_EFFECT_SIZE"
	// Stage Effect Rotation
	TextStageEffectRotation Text = "#STAGE_EFFECT_ROTATION"
	// Stage Effect Direction
	TextStageEffectDirection Text = "#STAGE_EFFECT_DIRECTION"
	// Stage Effect Transparency
	TextStageEffectAlpha Text = "#STAGE_EFFECT_ALPHA"
	// Lane
	TextLane Text = "#LANE"
	// Lane Position
	TextLanePosition Text = "#LANE_POSITION"
	// Lane Size
	TextLaneSize Text = "#LANE_SIZE"
	// Lane Rotation
	TextLaneRotation Text = "#LANE_ROTATION"
	// Lane Direction
	TextLaneDirection Text = "#LANE_DIRECTION"
	// Lane Transparency
	TextLaneAlpha Text = "#LANE_ALPHA"
	// Lane Animation
	TextLaneAnimation Text = "#LANE_ANIMATION"
	// Lane Effect
	TextLaneEffect Text = "#LANE_EFFECT"
	// Lane Effect Position
	TextLaneEffectPosition Text = "#LANE_EFFECT_POSITION"
	// Lane Effect Size
	TextLaneEffectSize Text = "#LANE_EFFECT_SIZE"
	// Lane Effect Rotation
	TextLaneEffectRotation Text = "#LANE_EFFECT_ROTATION"
	// Lane Effect Direction
	TextLaneEffectDirection Text = "#LANE_EFFECT_DIRECTION"
	// Lane Effect Transparency
	TextLaneEffectAlpha Text = "#LANE_EFFECT_ALPHA"
	// Judgment Line
	TextJudgeline Text = "#JUDGELINE"
	// Judgment Line Position
	TextJudgelinePosition Text = "#JUDGELINE_POSITION"
	// Judgment Line Size
	TextJudgelineSize Text = "#JUDGELINE_SIZE"
	// Judgment Line Rotation
	TextJudgelineRotation Text = "#JUDGELINE_ROTATION"
	// Judgment Line Direction
	TextJudgelineDirection Text = "#JUDGELINE_DIRECTION"
	// Judgment Line Transparency
	TextJudgelineAlpha Text = "#JUDGELINE_ALPHA"
	// Judgment Line Animation
	TextJudgelineAnimation Text = "#JUDGELINE_ANIMATION"
	// Judgment Line Effect
	TextJudgelineEffect Text = "#JUDGELINE_EFFECT"
	// Judgment Line Effect Position
	TextJudgelineEffectPosition Text = "#JUDGELINE_EFFECT_POSITION"
	// Judgment Line Effect Size
	TextJudgelineEffectSize Text = "#JUDGELINE_EFFECT_SIZE"
	// Judgment Line Effect Rotation
	TextJudgelineEffectRotation Text = "#JUDGELINE_EFFECT_ROTATION"
	// Judgment Line Effect Direction
	TextJudgelineEffectDirection Text = "#JUDGELINE_EFFECT_DIRECTION"
	// Judgment Line Effect Transparency
	TextJudgelineEffectAlpha Text = "#JUDGELINE_EFFECT_ALPHA"
	// Slot
	TextSlot Text = "#SLOT"
	// Slot Position
	TextSlotPosition Text = "#SLOT_POSITION"
	// Slot Size
	TextSlotSize Text = "#SLOT_SIZE"
	// Slot Rotation
	TextSlotRotation Text = "#SLOT_ROTATION"
	// Slot Direction
	TextSlotDirection Text = "#SLOT_DIRECTION"
	// Slot Transparency
	TextSlotAlpha Text = "#SLOT_ALPHA"
	// Slot Animation
	TextSlotAnimation Text = "#SLOT_ANIMATION"
	// Slot Effect
	TextSlotEffect Text = "#SLOT_EFFECT"
	// Slot Effect Position
	TextSlotEffectPosition Text = "#SLOT_EFFECT_POSITION"
	// Slot Effect Size
	TextSlotEffectSize Text = "#SLOT_EFFECT_SIZE"
	// Slot Effect Rotation
	TextSlotEffectRotation Text = "#SLOT_EFFECT_ROTATION"
	// Slot Effect Direction
	TextSlotEffectDirection Text = "#SLOT_EFFECT_DIRECTION"
	// Slot Effect Transparency
	TextSlotEffectAlpha Text = "#SLOT_EFFECT_ALPHA"
	// Note
	TextNote Text = "#NOTE"
	// Note Speed
	TextNoteSpeed Text = "#NOTE_SPEED"
	// Random Note Speed
	TextNoteSpeedRandom Text = "#NOTE_SPEED_RANDOM"
	// Note Position
	TextNotePosition Text = "#NOTE_POSITION"
	// Note Size
	TextNoteSize Text = "#NOTE_SIZE"
	// Note Rotation
	TextNoteRotation Text = "#NOTE_ROTATION"
	// Note Direction
	TextNoteDirection Text = "#NOTE_DIRECTION"
	// Note Color
	TextNoteColor Text = "#NOTE_COLOR"
	// Note Transparency
	TextNoteAlpha Text = "#NOTE_ALPHA"
	// Note Animation
	TextNoteAnimation Text = "#NOTE_ANIMATION"
	// Note Effect
	TextNoteEffect Text = "#NOTE_EFFECT"
	// Note Effect Position
	TextNoteEffectPosition Text = "#NOTE_EFFECT_POSITION"
	// Note Effect Size
	TextNoteEffectSize Text = "#NOTE_EFFECT_SIZE"
	// Note Effect Rotation
	TextNoteEffectRotation Text = "#NOTE_EFFECT_ROTATION"
	// Note Effect Direction
	TextNoteEffectDirection Text = "#NOTE_EFFECT_DIRECTION"
	// Note Effect Color
	TextNoteEffectColor Text = "#NOTE_EFFECT_COLOR"
	// Note Effect Transparency
	TextNoteEffectAlpha Text = "#NOTE_EFFECT_ALPHA"
	// Marker
	TextMarker Text = "#MARKER"
	// Marker Position
	TextMarkerPosition Text = "#MARKER_POSITION"
	// Marker Size
	TextMarkerSize Text = "#MARKER_SIZE"
	// Marker Rotation
	TextMarkerRotation Text = "#MARKER_ROTATION"
	// Marker Direction
	TextMarkerDirection Text = "#MARKER_DIRECTION"
	// Marker Color
	TextMarkerColor Text = "#MARKER_COLOR"
	// Marker Transparency
	TextMarkerAlpha Text = "#MARKER_ALPHA"
	// Marker Animation
	TextMarkerAnimation Text = "#MARKER_ANIMATION"
	// Connector
	TextConnector Text = "#CONNECTOR"
	// Connector Position
	TextConnectorPosition Text = "#CONNECTOR_POSITION"
	// Connector Size
	TextConnectorSize Text = "#CONNECTOR_SIZE"
	// Connector Rotation
	TextConnectorRotation Text = "#CONNECTOR_ROTATION"
	// Connector Direction
	TextConnectorDirection Text = "#CONNECTOR_DIRECTION"
	// Connector Color
	TextConnectorColor Text = "#CONNECTOR_COLOR"
	// Connector Transparency
	TextConnectorAlpha Text = "#CONNECTOR_ALPHA"
	// Connector Animation
	TextConnectorAnimation Text = "#CONNECTOR_ANIMATION"
	// Simultaneous Line
	TextSimline Text = "#SIMLINE"
	// Simultaneous Line Position
	TextSimlinePosition Text = "#SIMLINE_POSITION"
	// Simultaneous Line Size
	TextSimlineSize Text = "#SIMLINE_SIZE"
	// Simultaneous Line Rotation
	TextSimlineRotation Text = "#SIMLINE_ROTATION"
	// Simultaneous Line Direction
	TextSimlineDirection Text = "#SIMLINE_DIRECTION"
	// Simultaneous Line Color
	TextSimlineColor Text = "#SIMLINE_COLOR"
	// Simultaneous Line Transparency
	TextSimlineAlpha Text = "#SIMLINE_ALPHA"
	// Simultaneous Line Animation
	TextSimlineAnimation Text = "#SIMLINE_ANIMATION"
	// Preview Vertical Scale
	TextPreviewScaleVertical Text = "#PREVIEW_SCALE_VERTICAL"
	// Preview Horizontal Scale
	TextPreviewScaleHorizontal Text = "#PREVIEW_SCALE_HORIZONTAL"
	// Preview Time
	TextPreviewTime Text = "#PREVIEW_TIME"
	// Preview Score
	TextPreviewScore Text = "#PREVIEW_SCORE"
	// Preview BPM
	TextPreviewBpm Text = "#PREVIEW_BPM"
	// Preview Time Scale
	TextPreviewTimescale Text = "#PREVIEW_TIMESCALE"
	// Preview Beat
	TextPreviewBeat Text = "#PREVIEW_BEAT"
	// Preview Measure
	TextPreviewMeasure Text = "#PREVIEW_MEASURE"
	// Preview Combo
	TextPreviewCombo Text = "#PREVIEW_COMBO"
	// UI
	TextUi Text = "#UI"
	// UI Metric
	TextUiMetric Text = "#UI_METRIC"
	// UI Primary Metric
	TextUiPrimaryMetric Text = "#UI_PRIMARY_METRIC"
	// UI Secondary Metric
	TextUiSecondaryMetric Text = "#UI_SECONDARY_METRIC"
	// UI Judgment
	TextUiJudgment Text = "#UI_JUDGMENT"
	// UI Combo
	TextUiCombo Text = "#UI_Combo"
	// UI Menu
	TextUiMenu Text = "#UI_Menu"
	// ON
	TextOn Text = "#ON"
	// OFF
	TextOff Text = "#OFF"
	// None
	TextNone Text = "#NONE"
	// Any
	TextAny Text = "#ANY"
	// All
	TextAll Text = "#ALL"
	// Others
	TextOthers Text = "#OTHERS"
	// Short
	TextShort Text = "#SHORT"
	// Long
	TextLong Text = "#LONG"
	// High
	TextHigh Text = "#HIGH"
	// Mid
	TextMid Text = "#MID"
	// Low
	TextLow Text = "#LOW"
	// Small
	TextSmall Text = "#SMALL"
	// Medium
	TextMedium Text = "#MEDIUM"
	// Large
	TextLarge Text = "#LARGE"
	// Left
	TextLeft Text = "#LEFT"
	// Right
	TextRight Text = "#RIGHT"
	// Up
	TextUp Text = "#UP"
	// Down
	TextDown Text = "#DOWN"
	// Front
	TextFront Text = "#FRONT"
	// Back
	TextBack Text = "#BACK"
	// Center
	TextCenter Text = "#CENTER"
	// Top
	TextTop Text = "#TOP"
	// Bottom
	TextBottom Text = "#BOTTOM"
	// Top Left
	TextTopLeft Text = "#TOP_LEFT"
	// Top Center
	TextTopCenter Text = "#TOP_CENTER"
	// Top Right
	TextTopRight Text = "#TOP_RIGHT"
	// Center Left
	TextCenterLeft Text = "#CENTER_LEFT"
	// Center Right
	TextCenterRight Text = "#CENTER_RIGHT"
	// Bottom Left
	TextBottomLeft Text = "#BOTTOM_LEFT"
	// Bottom Center
	TextBottomCenter Text = "#BOTTOM_CENTER"
	// Bottom Right
	TextBottomRight Text = "#BOTTOM_RIGHT"
	// Clockwise
	TextClockwise Text = "#CLOCKWISE"
	// Counterclockwise
	TextCounterclockwise Text = "#COUNTERCLOCKWISE"
	// Forward
	TextForward Text = "#FORWARD"
	// Backward
	TextBackward Text = "#BACKWARD"
	// Default
	TextDefault Text = "#DEFAULT"
	// Neutral
	TextNeutral Text = "#NEUTRAL"
	// Red
	TextRed Text = "#RED"
	// Green
	TextGreen Text = "#GREEN"
	// Blue
	TextBlue Text = "#BLUE"
	// Yellow
	TextYellow Text = "#YELLOW"
	// Purple
	TextPurple Text = "#PURPLE"
	// Cyan
	TextCyan Text = "#CYAN"
	// Simple
	TextSimple Text = "#SIMPLE"
	// Easy
	TextEasy Text = "#EASY"
	// Normal
	TextNormal Text = "#NORMAL"
	// Hard
	TextHard Text = "#HARD"
	// Expert
	TextExpert Text = "#EXPERT"
	// Master
	TextMaster Text = "#MASTER"
	// Pro
	TextPro Text = "#PRO"
	// Technical
	TextTechnical Text = "#TECHNICAL"
	// Special
	TextSpecial Text = "#SPECIAL"
	// Append
	TextAppend Text = "#APPEND"
	// Enter post...
	TextPostPlaceholder Text = "#POST_PLACEHOLDER"
	// Enter playlist...
	TextPlaylistPlaceholder Text = "#PLAYLIST_PLACEHOLDER"
	// Enter level...
	TextLevelPlaceholder Text = "#LEVEL_PLACEHOLDER"
	// Enter skin...
	TextSkinPlaceholder Text = "#SKIN_PLACEHOLDER"
	// Enter background...
	TextBackgroundPlaceholder Text = "#BACKGROUND_PLACEHOLDER"
	// Enter SFX...
	TextEffectPlaceholder Text = "#EFFECT_PLACEHOLDER"
	// Enter particle...
	TextParticlePlaceholder Text = "#PARTICLE_PLACEHOLDER"
	// Enter engine...
	TextEnginePlaceholder Text = "#ENGINE_PLACEHOLDER"
	// Enter replay...
	TextReplayPlaceholder Text = "#REPLAY_PLACEHOLDER"
	// Enter user...
	TextUserPlaceholder Text = "#USER_PLACEHOLDER"
	// Enter room...
	TextRoomPlaceholder Text = "#ROOM_PLACEHOLDER"
	// Enter keywords...
	TextKeywordsPlaceholder Text = "#KEYWORDS_PLACEHOLDER"
	// Enter name...
	TextNamePlaceholder Text = "#NAME_PLACEHOLDER"
	// Enter rating...
	TextRatingPlaceholder Text = "#RATING_PLACEHOLDER"
	// Enter minimum rating...
	TextRatingMinimumPlaceholder Text = "#RATING_MINIMUM_PLACEHOLDER"
	// Enter maximum rating...
	TextRatingMaximumPlaceholder Text = "#RATING_MAXIMUM_PLACEHOLDER"
	// Enter title...
	TextTitlePlaceholder Text = "#TITLE_PLACEHOLDER"
	// Enter subtitle...
	TextSubtitlePlaceholder Text = "#SUBTITLE_PLACEHOLDER"
	// Enter artists...
	TextArtistsPlaceholder Text = "#ARTISTS_PLACEHOLDER"
	// Enter time...
	TextTimePlaceholder Text = "#TIME_PLACEHOLDER"
	// Enter author...
	TextAuthorPlaceholder Text = "#AUTHOR_PLACEHOLDER"
	// Enter coauthor...
	TextCoauthorPlaceholder Text = "#COAUTHOR_PLACEHOLDER"
	// Enter collaborator...
	TextCollaboratorPlaceholder Text = "#COLLABORATOR_PLACEHOLDER"
	// Enter description...
	TextDescriptionPlaceholder Text = "#DESCRIPTION_PLACEHOLDER"
	// Enter genre...
	TextGenrePlaceholder Text = "#GENRE_PLACEHOLDER"
	// Enter type...
	TextTypePlaceholder Text = "#TYPE_PLACEHOLDER"
	// Enter category...
	TextCategoryPlaceholder Text = "#CATEGORY_PLACEHOLDER"
	// Enter language...
	TextLanguagePlaceholder Text = "#LANGUAGE_PLACEHOLDER"
	// Enter difficulty...
	TextDifficultyPlaceholder Text = "#DIFFICULTY_PLACEHOLDER"
	// Enter length...
	TextLengthPlaceholder Text = "#LENGTH_PLACEHOLDER"
	// Enter minimum length...
	TextLengthMinimumPlaceholder Text = "#LENGTH_MINIMUM_PLACEHOLDER"
	// Enter maximum length...
	TextLengthMaximumPlaceholder Text = "#LENGTH_MAXIMUM_PLACEHOLDER"
	// Enter additional information...
	TextAdditionalInformationPlaceholder Text = "#ADDITIONAL_INFORMATION_PLACEHOLDER"
	// Enter timezone...
	TextTimezonePlaceholder Text = "#TIMEZONE_PLACEHOLDER"
	// Enter region...
	TextRegionPlaceholder Text = "#REGION_PLACEHOLDER"
	// Enter content...
	TextContentPlaceholder Text = "#CONTENT_PLACEHOLDER"
	// Enter comment...
	TextCommentPlaceholder Text = "#COMMENT_PLACEHOLDER"
	// Enter review...
	TextReviewPlaceholder Text = "#REVIEW_PLACEHOLDER"
	// Enter reply...
	TextReplyPlaceholder Text = "#REPLY_PLACEHOLDER"
	// Enter message...
	TextMessagePlaceholder Text = "#MESSAGE_PLACEHOLDER"
	// Enter role...
	TextRolePlaceholder Text = "#ROLE_PLACEHOLDER"
	// Enter permission...
	TextPermissionPlaceholder Text = "#PERMISSION_PLACEHOLDER"
	// {0}%
	TextPercentageUnit Text = "#PERCENTAGE_UNIT"
	// {0}yr
	TextYearUnit Text = "#YEAR_UNIT"
	// {0}mo
	TextMonthUnit Text = "#MONTH_UNIT"
	// {0}d
	TextDayUnit Text = "#DAY_UNIT"
	// {0}h
	TextHourUnit Text = "#HOUR_UNIT"
	// {0}m
	TextMinuteUnit Text = "#MINUTE_UNIT"
	// {0}s
	TextSecondUnit Text = "#SECOND_UNIT"
	// {0}ms
	TextMillisecondUnit Text = "#MILLISECOND_UNIT"
	// {0}yr ago
	TextYearPast Text = "#YEAR_PAST"
	// {0}mo ago
	TextMonthPast Text = "#MONTH_PAST"
	// {0}d ago
	TextDayPast Text = "#DAY_PAST"
	// {0}h ago
	TextHourPast Text = "#HOUR_PAST"
	// {0}m ago
	TextMinutePast Text = "#MINUTE_PAST"
	// {0}s ago
	TextSecondPast Text = "#SECOND_PAST"
	// {0}ms ago
	TextMillisecondPast Text = "#MILLISECOND_PAST"
	// In {0}yr
	TextYearFuture Text = "#YEAR_FUTURE"
	// In {0}mo
	TextMonthFuture Text = "#MONTH_FUTURE"
	// In {0}d
	TextDayFuture Text = "#DAY_FUTURE"
	// In {0}h
	TextHourFuture Text = "#HOUR_FUTURE"
	// In {0}m
	TextMinuteFuture Text = "#MINUTE_FUTURE"
	// In {0}s
	TextSecondFuture Text = "#SECOND_FUTURE"
	// In {0}ms
	TextMillisecondFuture Text = "#MILLISECOND_FUTURE"
	// Tap
	TextTap Text = "#TAP"
	// Tap and Hold
	TextTapHold Text = "#TAP_HOLD"
	// Tap and Release
	TextTapRelease Text = "#TAP_RELEASE"
	// Tap and Flick
	TextTapFlick Text = "#TAP_FLICK"
	// Tap and Slide
	TextTapSlide Text = "#TAP_SLIDE"
	// Hold
	TextHold Text = "#HOLD"
	// Hold and Slide
	TextHoldSlide Text = "#HOLD_SLIDE"
	// Hold and Follow
	TextHoldFollow Text = "#HOLD_FOLLOW"
	// Release
	TextRelease Text = "#RELEASE"
	// Flick
	TextFlick Text = "#FLICK"
	// Slide
	TextSlide Text = "#SLIDE"
	// Slide and Flick
	TextSlideFlick Text = "#SLIDE_FLICK"
	// Avoid
	TextAvoid Text = "#AVOID"
	// Jiggle
	TextJiggle Text = "#JIGGLE"
	// Newest
	TextNewest Text = "#NEWEST"
	// Oldest
	TextOldest Text = "#OLDEST"
	// Recommended
	TextRecommended Text = "#RECOMMENDED"
	// Popular
	TextPopular Text = "#POPULAR"
	// Featured
	TextFeatured Text = "#FEATURED"
	// Competitive
	TextCompetitive Text = "#COMPETITIVE"
	// Tournament
	TextTournament Text = "#TOURNAMENT"
	// Holiday
	TextHoliday Text = "#HOLIDAY"
	// Limited
	TextLimited Text = "#LIMITED"
	// Announcement
	TextAnnouncement Text = "#ANNOUNCEMENT"
	// Information
	TextInformation Text = "#INFORMATION"
	// Help
	TextHelp Text = "#HELP"
	// Maintenance
	TextMaintenance Text = "#MAINTENANCE"
	// Event
	TextEvent Text = "#EVENT"
	// Update
	TextUpdate Text = "#UPDATE"
	// Search
	TextSearch Text = "#SEARCH"
	// Advanced
	TextAdvanced Text = "#ADVANCED"
	// Related
	TextRelated Text = "#RELATED"
	// Same Author
	TextSameAuthor Text = "#SAME_AUTHOR"
	// Same Artists
	TextSameArtists Text = "#SAME_ARTISTS"
	// Same Rating
	TextSameRating Text = "#SAME_RATING"
	// Same Category
	TextSameCategory Text = "#SAME_CATEGORY"
	// Same Difficulty
	TextSameDifficulty Text = "#SAME_DIFFICULTY"
	// Same Genre
	TextSameGenre Text = "#SAME_GENRE"
	// Same Version
	TextSameVersion Text = "#SAME_VERSION"
	// Other Authors
	TextOtherAuthors Text = "#OTHER_AUTHORS"
	// Other Artists
	TextOtherArtists Text = "#OTHER_ARTISTS"
	// Other Ratings
	TextOtherRatings Text = "#OTHER_RATINGS"
	// Other Categories
	TextOtherCategories Text = "#OTHER_CATEGORIES"
	// Other Difficulties
	TextOtherDifficulties Text = "#OTHER_DIFFICULTIES"
	// Other Genres
	TextOtherGenres Text = "#OTHER_GENRES"
	// Other Versions
	TextOtherVersions Text = "#OTHER_VERSIONS"
	// Draft
	TextDraft Text = "#DRAFT"
	// Public
	TextPublic Text = "#PUBLIC"
	// Private
	TextPrivate Text = "#PRIVATE"
	// Pop
	TextPop Text = "#POP"
	// Rock
	TextRock Text = "#ROCK"
	// Hip Hop
	TextHiphop Text = "#HIPHOP"
	// Country
	TextCountry Text = "#COUNTRY"
	// Electronic
	TextElectronic Text = "#ELECTRONIC"
	// Metal
	TextMetal Text = "#METAL"
	// Classical
	TextClassical Text = "#CLASSICAL"
	// Folk
	TextFolk Text = "#FOLK"
	// Indie
	TextIndie Text = "#INDIE"
	// Anime
	TextAnime Text = "#ANIME"
	// Vocaloid
	TextVocaloid Text = "#VOCALOID"
	// Remix
	TextRemix Text = "#REMIX"
	// Instrumental
	TextInstrumental Text = "#INSTRUMENTAL"
	// Short Version
	TextShortVersion Text = "#SHORT_VERSION"
	// Long Version
	TextLongVersion Text = "#LONG_VERSION"
	// Cut Version
	TextCutVersion Text = "#CUT_VERSION"
	// Full Version
	TextFullVersion Text = "#FULL_VERSION"
	// Extended Version
	TextExtendedVersion Text = "#EXTENDED_VERSION"
	// Live Version
	TextLiveVersion Text = "#LIVE_VERSION"
	// Medley
	TextMedley Text = "#MEDLEY"
	// Explicit
	TextExplicit Text = "#EXPLICIT"
	// Multi Finger
	TextMultiFinger Text = "#MULTI_FINGER"
	// Full Hand
	TextFullHand Text = "#FULL_HAND"
	// Cross Hand
	TextCrossHand Text = "#CROSS_HAND"
	// Gimmick
	TextGimmick Text = "#GIMMICK"
	// Collaboration
	TextCollaboration Text = "#COLLABORATION"
	// Report
	TextReport Text = "#REPORT"
	// Reason
	TextReason Text = "#REASON"
	// Illegal Activities
	TextIllegalActivities Text = "#ILLEGAL_ACTIVITIES"
	// Cheating
	TextCheating Text = "#CHEATING"
	// AFK
	TextAfk Text = "#AFK"
	// Spamming
	TextSpamming Text = "#SPAMMING"
	// Verbal Abuse
	TextVerbalAbuse Text = "#VERBAL_ABUSE"
	// Inappropriate Language
	TextInappropriateLanguage Text = "#INAPPROPRIATE_LANGUAGE"
	// Negative Attitude
	TextNegativeAttitude Text = "#NEGATIVE_ATTITUDE"
	// DNF
	TextDnf Text = "#DNF"
	// Suggestions
	TextSuggestions Text = "#SUGGESTIONS"
	// Suggestions per Player
	TextSuggestionsPerPlayer Text = "#SUGGESTIONS_PER_PLAYER"
	// Match Scoring
	TextMatchScoring Text = "#MATCH_SCORING"
	// Match Tiebreaker
	TextMatchTiebreaker Text = "#MATCH_TIEBREAKER"
	// Match Count
	TextMatchCount Text = "#MATCH_COUNT"
	// Match Limit
	TextMatchLimit Text = "#MATCH_LIMIT"
	// Round Scoring
	TextRoundScoring Text = "#ROUND_SCORING"
	// Round Tiebreaker
	TextRoundTiebreaker Text = "#ROUND_TIEBREAKER"
	// Round Count
	TextRoundCount Text = "#ROUND_COUNT"
	// Round Limit
	TextRoundLimit Text = "#ROUND_LIMIT"
	// Team Scoring
	TextTeamScoring Text = "#TEAM_SCORING"
	// Team Tiebreaker
	TextTeamTiebreaker Text = "#TEAM_TIEBREAKER"
	// Team Count
	TextTeamCount Text = "#TEAM_COUNT"
	// Team Limit
	TextTeamLimit Text = "#TEAM_LIMIT"
	// Qualified
	TextQualified Text = "#QUALIFIED"
	// Disqualified
	TextDisqualified Text = "#DISQUALIFIED"
	// Ranking
	TextRanking Text = "#RANKING"
	// Score
	TextScore Text = "#SCORE"
	// Owner
	TextOwner Text = "#OWNER"
	// Admin
	TextAdmin Text = "#ADMIN"
	// Moderator
	TextModerator Text = "#MODERATOR"
	// Reviewer
	TextReviewer Text = "#REVIEWER"
	// Banned
	TextBanned Text = "#BANNED"
	// Player
	TextPlayer Text = "#PLAYER"
	// Spectator
	TextSpectator Text = "#SPECTATOR"
	// Referee
	TextReferee Text = "#REFEREE"
	// Eliminated
	TextEliminated Text = "#ELIMINATED"
	// Finalist
	TextFinalist Text = "#FINALIST"
	// Finished
	TextFinished Text = "#FINISHED"
	// Winner
	TextWinner Text = "#WINNER"
	// Gold Medal
	TextGoldMedal Text = "#GOLD_MEDAL"
	// Silver Medal
	TextSilverMedal Text = "#SILVER_MEDAL"
	// Bronze Medal
	TextBronzeMedal Text = "#BRONZE_MEDAL"
	// Team 1
	TextTeam1 Text = "#TEAM_1"
	// Team 2
	TextTeam2 Text = "#TEAM_2"
	// Team 3
	TextTeam3 Text = "#TEAM_3"
	// Team 4
	TextTeam4 Text = "#TEAM_4"
	// Team 5
	TextTeam5 Text = "#TEAM_5"
	// Team 6
	TextTeam6 Text = "#TEAM_6"
	// Team 7
	TextTeam7 Text = "#TEAM_7"
	// Team 8
	TextTeam8 Text = "#TEAM_8"
	// Team Red
	TextTeamRed Text = "#TEAM_RED"
	// Team Green
	TextTeamGreen Text = "#TEAM_GREEN"
	// Team Blue
	TextTeamBlue Text = "#TEAM_BLUE"
	// Team Yellow
	TextTeamYellow Text = "#TEAM_YELLOW"
	// Team Purple
	TextTeamPurple Text = "#TEAM_PURPLE"
	// Team Cyan
	TextTeamCyan Text = "#TEAM_CYAN"
	// Team White
	TextTeamWhite Text = "#TEAM_WHITE"
	// Team Black
	TextTeamBlack Text = "#TEAM_BLACK"
	// Add
	TextAdd Text = "#ADD"
	// Added
	TextAdded Text = "#ADDED"
	// Create
	TextCreate Text = "#CREATE"
	// Created
	TextCreated Text = "#CREATED"
	// Reply
	TextReply Text = "#REPLY"
	// Replied
	TextReplied Text = "#REPLIED"
	// Review
	TextReview Text = "#REVIEW"
	// Reviewing
	TextReviewing Text = "#REVIEWING"
	// Reviewed
	TextReviewed Text = "#REVIEWED"
	// Verify
	TextVerify Text = "#VERIFY"
	// Verifying
	TextVerifying Text = "#VERIFYING"
	// Verified
	TextVerified Text = "#VERIFIED"
	// Upload
	TextUpload Text = "#UPLOAD"
	// Uploading
	TextUploading Text = "#UPLOADING"
	// Uploaded
	TextUploaded Text = "#UPLOADED"
	// Submit
	TextSubmit Text = "#SUBMIT"
	// Submitting
	TextSubmitting Text = "#SUBMITTING"
	// Submitted
	TextSubmitted Text = "#SUBMITTED"
	// Edit
	TextEdit Text = "#EDIT"
	// Editing
	TextEditing Text = "#EDITING"
	// Edited
	TextEdited Text = "#EDITED"
	// Like
	TextLike Text = "#LIKE"
	// Liked
	TextLiked Text = "#LIKED"
	// Dislike
	TextDislike Text = "#DISLIKE"
	// Disliked
	TextDisliked Text = "#DISLIKED"
	// Bookmark
	TextBookmark Text = "#BOOKMARK"
	// Bookmarked
	TextBookmarked Text = "#BOOKMARKED"
	// Delete
	TextDelete Text = "#DELETE"
	// Deleting
	TextDeleting Text = "#DELETING"
	// Deleted
	TextDeleted Text = "#DELETED"
	// Remove
	TextRemove Text = "#REMOVE"
	// Removing
	TextRemoving Text = "#REMOVING"
	// Removed
	TextRemoved Text = "#REMOVED"
	// Restore
	TextRestore Text = "#RESTORE"
	// Restoring
	TextRestoring Text = "#RESTORING"
	// Restored
	TextRestored Text = "#RESTORED"
	// Confirm
	TextConfirm Text = "#CONFIRM"
	// Confirmed
	TextConfirmed Text = "#CONFIRMED"
	// Cancel
	TextCancel Text = "#CANCEL"
	// Canceled
	TextCanceled Text = "#CANCELED"
	// Increase
	TextIncrease Text = "#INCREASE"
	// Decrease
	TextDecrease Text = "#DECREASE"
	// Upvote
	TextUpvote Text = "#UPVOTE"
	// Upvoted
	TextUpvoted Text = "#UPVOTED"
	// Downvote
	TextDownvote Text = "#DOWNVOTE"
	// Downvoted
	TextDownvoted Text = "#DOWNVOTED"
	// Agree
	TextAgree Text = "#AGREE"
	// Agreed
	TextAgreed Text = "#AGREED"
	// Disagree
	TextDisagree Text = "#DISAGREE"
	// Disagreed
	TextDisagreed Text = "#DISAGREED"
	// Lock
	TextLock Text = "#LOCK"
	// Locked
	TextLocked Text = "#LOCKED"
	// Unlock
	TextUnlock Text = "#UNLOCK"
	// Unlocked
	TextUnlocked Text = "#UNLOCKED"
	// Pin
	TextPin Text = "#PIN"
	// Pinned
	TextPinned Text = "#PINNED"
	// Unpin
	TextUnpin Text = "#UNPIN"
	// Unpinned
	TextUnpinned Text = "#UNPINNED"
	// Follow
	TextFollow Text = "#FOLLOW"
	// Following
	TextFollowing Text = "#FOLLOWING"
	// Followed
	TextFollowed Text = "#FOLLOWED"
	// Unfollow
	TextUnfollow Text = "#UNFOLLOW"
	// Subscribe
	TextSubscribe Text = "#SUBSCRIBE"
	// Subscribing
	TextSubscribing Text = "#SUBSCRIBING"
	// Subscribed
	TextSubscribed Text = "#SUBSCRIBED"
	// Unsubscribe
	TextUnsubscribe Text = "#UNSUBSCRIBE"
	// Publish
	TextPublish Text = "#PUBLISH"
	// Publishing
	TextPublishing Text = "#PUBLISHING"
	// Published
	TextPublished Text = "#PUBLISHED"
	// Unpublish
	TextUnpublish Text = "#UNPUBLISH"
	// Show
	TextShow Text = "#SHOW"
	// Hide
	TextHide Text = "#HIDE"
	// Allow
	TextAllow Text = "#ALLOW"
	// Allowed
	TextAllowed Text = "#ALLOWED"
	// Disallow
	TextDisallow Text = "#DISALLOW"
	// Disallowed
	TextDisallowed Text = "#DISALLOWED"
	// Approve
	TextApprove Text = "#APPROVE"
	// Approved
	TextApproved Text = "#APPROVED"
	// Deny
	TextDeny Text = "#DENY"
	// Denied
	TextDenied Text = "#DENIED"
	// Accept
	TextAccept Text = "#ACCEPT"
	// Accepted
	TextAccepted Text = "#ACCEPTED"
	// Reject
	TextReject Text = "#REJECT"
	// Rejected
	TextRejected Text = "#REJECTED"
	// Star
	TextStar Text = "#STAR"
	// Starred
	TextStarred Text = "#STARRED"
)
