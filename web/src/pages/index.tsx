import type { NextPage } from 'next'
import Head from 'next/head'
import { Heading, Box, Divider, VStack } from '@chakra-ui/react'
import HomeMain from '../components/HomeMain'

const Home: NextPage = () => {
  return (
    <>
      <Box py="100px" width="100vw" backgroundImage="https://images.unsplash.com/photo-1556519272-32c24ecd2253?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1074&q=80">
        <Heading textAlign="center">
          Yzlhdy的博客
        </Heading >
        <Heading as="h6" textAlign="center" fontSize="14" my="10">
          为制造商、开发人员和设计师收集免费的设计工具和资源
        </Heading>
      </Box>
      <Box w="container.xl" >
        <Divider />
        <HomeMain />
      </Box>
    </>
  )
}

export default Home
