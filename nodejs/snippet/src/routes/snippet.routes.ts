import { Router } from "express";
import { createSnippet, getSnippet } from "../controllers/snippet.controller";
const snippetRouter = Router();

snippetRouter.post("/", createSnippet).get("/", getSnippet);

export default snippetRouter;
