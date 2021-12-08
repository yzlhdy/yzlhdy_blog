import React from 'react';
import { Box, Grid, GridItem, VStack, Avatar, Heading, Stack, Button, Link } from '@chakra-ui/react';
import { GoOctoface } from 'react-icons/go'


const HomeMain: React.FC = () => {
	return (
		<Box spacing={0} w="full" mt={5} >
			<Grid
				templateColumns="repeat(6, 1fr)"
				gap={2}
			>
				<GridItem colSpan={2} bg='tomato' rounded="xl" boxShadow="xl">
					<VStack py={10}  >
						<Avatar size='2xl' name='Segun Adebayo' src='https://images.unsplash.com/photo-1511044568932-338cba0ad803?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8Y2F0fGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60' />
						<Heading py={4}>Yzlhdy</Heading>
						<Heading as="p" fontSize="14px">闻道有先后,术业有专攻,如是而已</Heading>
					</VStack>
					<Stack direction='row' spacing={4}>
						<Box display="flex" alignItems="center" justifyContent="center" w="full">
							<GoOctoface size={24} color="pink" />
							<Button as={Link} variant="link">https://github.com/yzlhdy</Button>
						</Box>
					</Stack>
				</GridItem>
				<GridItem colSpan={4} bg='papayawhip' >2</GridItem>
			</Grid>
		</Box>
	);
}

export default HomeMain;