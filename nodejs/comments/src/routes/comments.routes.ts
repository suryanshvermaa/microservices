import {Router} from 'express';
import { createComment, getCommentBySnippetId } from '../controllers/comments.controller';
const commentsRouter=Router();

commentsRouter
.post('/',createComment)
.get('/',getCommentBySnippetId)

export default commentsRouter;