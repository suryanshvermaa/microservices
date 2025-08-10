import { Router } from "express";
import snippetRouter from "./snippet.routes";
const router = Router();

router.use("/snippet", snippetRouter);

export default router;
