import React from 'react';
import { useColorMode } from '@chakra-ui/react'
import { Button, IconButton } from '@chakra-ui/button';
import { Heading, HStack, Link, LinkBox, Text } from '@chakra-ui/layout';
import { IoMoonSharp, IoSunny } from 'react-icons/io5'
import NextLink from 'next/link'
import logoImage from '../../assets/images/logo.png'
import Image from 'next/image';

const Header: React.FC = () => {
	const { toggleColorMode, colorMode } = useColorMode();

	return (
		<HStack justifyContent="space-between" alignItems="center" py={4} as="nav" width="container.xl" >
			<LinkBox href="/" >
				<Image
					src={logoImage}
					alt="logo"
					width="40px"
					height="40px"
				/>
			</LinkBox>
			<HStack alignItems="center" spacing={10} >
				<NextLink href="/home" passHref>
					<Button as={Link} size="sm" variant="ghost">
						<Text
							_hover={{
								textDecoration: 'none'
							}}
						> 主页</Text>
					</Button>
				</NextLink>

				<NextLink href="/about" passHref>
					<Button as={Link} size="sm" variant="ghost">
						<Text
							_hover={{
								textDecoration: 'none'
							}}
						> 资源</Text>
					</Button>
				</NextLink>

				<NextLink href="/home" passHref>
					<Button as={Link} size="sm" variant="ghost">
						<Text
							_hover={{
								textDecoration: 'none'
							}}
						>分类</Text>
					</Button>
				</NextLink>
				<NextLink href="/home" passHref>
					<Button as={Link} size="sm" variant="ghost">
						<Text
							_hover={{
								textDecoration: 'none'
							}}
						>我的</Text>
					</Button>
				</NextLink>
				<IconButton
					aria-label="toggle theme"
					icon={colorMode === 'light' ? <IoMoonSharp /> : <IoSunny />}
					onClick={toggleColorMode}
				/>
			</HStack>
		</HStack >
	);
}

export default Header;