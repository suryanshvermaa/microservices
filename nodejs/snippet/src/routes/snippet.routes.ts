import {Router} from 'express';
import { createSnippet } from '../controllers/snippet.controller';
const snippetRouter=Router();

snippetRouter
.post('/',createSnippet)

export default snippetRouter;