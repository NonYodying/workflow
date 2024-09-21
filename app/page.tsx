'use client'

import React, { useEffect, useState } from 'react'
import axios from 'axios'


const List = () => {
  const [posts, setPosts] = useState([])

  useEffect(() => {
    fetchPosts()
  }, [])

  const fetchPosts = async () => {
    try {
      const res = await axios.get('/api/posts')
      setPosts(res.data)
    } catch (error) {
      console.error(error)
    }
  }

  const deletePost = async (id: Number) => {
    try {
      await axios.delete(`/api/posts/${id}`)
      fetchPosts()
    } catch (error) {
      console.error('Failed to delete the post', error)
    }
  }

  return (
    <div>SOME BEAUTIFUL DASHBOAD</div>
  )
}

export default List