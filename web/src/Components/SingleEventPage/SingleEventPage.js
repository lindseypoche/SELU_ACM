import React, { Component } from 'react';
import './SingleEventPage.css';
import { Sidebar, Nav, Button, Avatar, Image, Box } from 'grommet';
import { FaRegComment, FaRegHeart, FaGrinTongueWink } from 'react-icons/fa';
class SingleEventPage extends Component {

    render() {
        return (
            <div className="page-container">
                <div className="single-container">
                    <div className="title-container">
                        <h1>Doing something for some reason coming soon</h1>
                    </div>
                    <div className="details-container">
                        <Avatar
                            size="medium"
                            background="lightgray"
                            margin="small"
                        >
                            <FaGrinTongueWink />
                        </Avatar>
                        <Button
                            background="none"
                            color="green"
                            font-size="small"
                            margin="small"
                        >Ian Porter</Button>
                        <p className="text-issue">March 26, 2021</p>
                    </div>
                    <div className="img-container">
                        <Box
                            height="medium"
                            width="medium"
                        >
                            <Image
                                fit="contain"
                                src="https://via.placeholder.com/150"
                                fill="horizontal"
                                size="xxsmall"
                            />
                        </Box>
                    </div>
                    <div className="content-container">
                        <div className="left-side">
                            <Sidebar
                                background="brand"
                                round="small"
                                width="fit-content"
                                height="fit-content"
                                background={{ "color": "#2c2c2c" }}
                            >
                                <Nav gap="small">

                                    <Button icon={<FaRegHeart />} hoverIndicator />
                                    <Button icon={<FaRegComment />} hoverIndicator />
                                </Nav>
                            </Sidebar>
                        </div>

                        <div className="body-container">
                            <p className="font-mgmt">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
                        </div>
                    </div>
                </div >
            </div>
        );
    }
}
export default SingleEventPage;