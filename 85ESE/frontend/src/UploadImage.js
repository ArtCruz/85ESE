import React from 'react';
import Form from 'react-bootstrap/Form';
import Col from 'react-bootstrap/Col';
import Row from 'react-bootstrap/Row';
import Button from 'react-bootstrap/Button';
import Container from 'react-bootstrap/Container';
import Toast from './Toast.js';
import axios from 'axios';

class UploadImage extends React.Component {
    constructor(props) {
        super(props);
        // Pega o id da URL (React Router v5)
        const { id } = this.props.params;
        this.state = {
            validated: false,
            id: id || "",
            buttonDisabled: false,
            toastShow: false,
            toastText: ""
        };

        this.validated = this.validated.bind(this);
        this.changeHandler = this.changeHandler.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    validated() {
        return this.state.validated;
    }

    handleSubmit(event) {
        event.preventDefault();

        const data = new FormData();
        data.append('file', this.state.file);
        data.append('id', this.state.id);

        axios.post(window.global.gateway_location + '/images', data, {
            headers: {
                'Content-Type': 'multipart/form-data',
            },
        })
            .then((response) => {
                this.setState({ toastShow: true, toastText: "Imagem enviada com sucesso!" });
            })
            .catch((error) => {
                this.setState({ toastShow: true, toastText: "Erro ao enviar imagem." });
                console.error("Erro ao enviar imagem:", error);
            });
    }

    changeHandler(event) {
        if (event.target.name === "file") {
            this.setState({ [event.target.name]: event.target.files[0], toastShow: false });
            return;
        }
        this.setState({ [event.target.name]: event.target.value, toastShow: false });
    }

    render() {
        return (
            <div>
                <h1 style={{ marginBottom: "40px" }}>Upload de Imagem</h1>
                <Container className="text-left">
                    <Form noValidate validated={this.validated} onSubmit={this.handleSubmit}>
                        <Form.Group as={Row} controlId="productID">
                            <Form.Label column sm="2">Product ID:</Form.Label>
                            <Col sm="6">
                                <Form.Control
                                    type="text"
                                    name="id"
                                    value={this.state.id}
                                    disabled
                                    style={{ width: "80px" }}
                                />
                                <Form.Text className="text-muted">
                                    ID do produto para associar a imagem
                                </Form.Text>
                            </Col>
                            <Col sm="4">
                                <Toast show={this.state.toastShow} message={this.state.toastText} />
                            </Col>
                        </Form.Group>
                        <Form.Group as={Row}>
                            <Form.Label column sm="2">File:</Form.Label>
                            <Col sm="10">
                                <Form.Control type="file" name="file" required onChange={this.changeHandler} />
                                <Form.Text className="text-muted">Imagem para associar ao produto</Form.Text>
                                <Form.Control.Feedback type="invalid">Selecione um arquivo para upload.</Form.Control.Feedback>
                            </Col>
                        </Form.Group>
                        <Button type="submit" disabled={this.state.buttonDisabled}>Enviar</Button>
                    </Form>
                </Container>
            </div>
        );
    }
}

export default UploadImage;