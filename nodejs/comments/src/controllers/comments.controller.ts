import asyncHandler from "../utils/asyncHandler";
import { Request,Response } from "express";

/**
* @description Create a new comment
* @route POST /api/v1/comment
* @access Public
* @param {Request} req
* @param {Response} res
*/
export const createComment = asyncHandler(async (req: Request, res: Response) => {
    
});

/**
* @description Get a comment By snippetId
* @route POST /api/v1/comment
* @access Public
* @param {Request} req
* @param {Response} res
*/
export const getCommentBySnippetId = asyncHandler(async (req: Request, res: Response) => {
    
});