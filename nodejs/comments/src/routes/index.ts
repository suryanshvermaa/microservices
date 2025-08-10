import { Router } from "express";
import commentsRouter from "./comments.routes";
const router = Router();

router.use("/comment", commentsRouter);

export default router;
