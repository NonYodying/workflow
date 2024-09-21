'use client'

import React, { useEffect, useState } from 'react'
import axios from 'axios'


import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useRouter } from 'next/navigation'

const Edit = ({ params }: { params: { id: string } }) => {
    const [title, setTitle] = useState('')
    const [status, setStatus] = useState('')
    const router = useRouter()
    const { id } = params

    const fetchPost = async (id: Number) => {
        try {
            const res = await axios.get(`/api/posts/${id}`)
            setStatus(res.data.status)
            setTitle(res.data.title)

        } catch (error) {
            console.error(error)
        }
    }

    useEffect(() => {
        if (id) {
            fetchPost(parseInt(id))
        }
    }, [id])

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()

        try {
            await axios.patch(`/api/posts/${id}`, { status })
            router.push('/entry')
        } catch (error) {
            console.error(error)
        }
    }

    const handleCancle = async (e: React.FormEvent) => {
        e.preventDefault()
        router.push('/entry')
    }

    return (
        <Card className="w-[350px]">
            <CardHeader>
                <CardTitle>Approve or Reject {id}</CardTitle>
                <CardDescription>Enter APPROVED or REJECTED ONLY</CardDescription>

            </CardHeader>
            <CardContent>
                <form>
                    <div className="grid w-full items-center gap-4">
                        <CardTitle>Title : {title}</CardTitle>
                        <div className="flex flex-col space-y-1.5">
                            <Label htmlFor="title">Status</Label>
                            <Input
                                id="status"
                                type="text"
                                name="status"
                                required
                                value={status}
                                onChange={(e) => setStatus(e.target.value)} />
                        </div>

                    </div>
                </form>
            </CardContent>
            <CardFooter className="flex justify-between">
                <Button onClick={handleCancle} variant="outline">Cancel</Button>
                <Button onClick={handleSubmit}>Edit</Button>
            </CardFooter>
        </Card>
    )
}

export default Edit