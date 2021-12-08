import { Container, VStack } from '@chakra-ui/layout';

import React, { PropsWithChildren } from 'react';
import Footer from './Footer';
import Header from './Header';
type Props = PropsWithChildren<{}>;

const Layout: React.FC<Props> = ({ children }) => {
	return (
		<Container>
			<VStack spacing={0} >
				<Header />
				{
					children
				}
				<Footer />

			</VStack>
		</Container>

	)
}

export default Layout;