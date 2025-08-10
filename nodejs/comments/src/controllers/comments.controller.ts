import { comments, IComment } from "../database/index.database";
import asyncHandler from "../utils/asyncHandler";
import { Request, Response } from "express";
import response from "../utils/response";
import { AppError } from "../utils/error";

/**
 * @description Create a new comment
 * @route POST /api/v1/comment
 * @access Public
 * @param {Request} req
 * @param {Response} res
 */
export const createComment = asyncHandler(
	async (req: Request, res: Response) => {
		const id = Date.now().toString();
		const body: IComment = req.body as unknown as IComment;
		body.id = id;
		comments.push(body);
		response(res, 200, "success", body);
	}
);

/**
 * @description Get a comment By snippetId
 * @route POST /api/v1/comment/:id
 * @access Public
 * @param {Request} req
 * @param {Response} res
 */
export const getCommentBySnippetId = asyncHandler(
	async (req: Request, res: Response) => {
		const { id } = req.params;
		const cmts = comments.find((c) => c.snippet_id === id);
		if (!cmts) throw new AppError("Comment not found", 404);
		response(res, 200, "success", cmts);
	}
);
