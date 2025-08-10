import React from 'react'
import { Input } from './ui/input'
import { Textarea } from './ui/textarea'
import { Button } from './ui/button'

const CreateSnippet: React.FC = () => {
  return (
    <div>
        <form className='mt-10 flex flex-col gap-4'>
            <Input type='text' placeholder='write a code snippets...'/>
            <Textarea placeholder='write a code snippets...'/>
            <div className='flex justify-start'>
                <Button>Add</Button>
            </div>
        </form>
    </div>
  )
}

export default CreateSnippet