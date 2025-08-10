import { Router } from "express";
import commentsRouter from "./comments.routes";
const router = Router();

router.use("/api/v1", commentsRouter);

export default router;
