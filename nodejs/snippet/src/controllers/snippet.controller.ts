import { randomBytes } from "crypto";
import asyncHandler from "../utils/asyncHandler";
import { ISnippet, snippet } from "../database/index.database";
import { AppError } from "../utils/error";
import { Request,Response } from "express";
import response from "../utils/response";

/**
* @description Create a new snippet
* @route POST /api/v1/snippet
* @access Public
* @param {Request} req
* @param {Response} res
*/
export const createSnippet = asyncHandler(async (req: Request, res: Response) => {
   const id=randomBytes(4).toString('hex');
   const body:ISnippet=req.body as any;
   if(!body.title || !body.description || !body.code || !body.userId){
      throw new AppError("All fields are required",400);
   }
   body.id=id;
   snippet.push(body);
   console.log(snippet)
   response(res,200,"Snippet created successfully",body);
});