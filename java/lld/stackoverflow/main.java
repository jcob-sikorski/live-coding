// UserRegistry is going to be a singleton
//
// i'd guess when it comes to timeline we are going to just query the database for most recent posts
// 
// for each user posting, answering and commenting on timeline i'd use a command pattern
// 
// there's an association of question with a tag
// 
// there's an association between questions, (many to one) with user profiles
// 
// i think searching for user profiles, tags, keywords could be implemented using inverted index
// 
// there's going to be used an observer pattern, when user posts a question the system is going to observe the post added to the post registry  and qualify fhe contribution of the user and assign reputation to that user + update his activity.
// 
// the system should handle concurrent access to the timeline
// 
// question-answer-comment graph attached to each question, state of each question (votes to the questions), that would be a composite pattern