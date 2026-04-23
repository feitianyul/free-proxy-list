**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-04-23 12:50:37 UTC（2026-04-23 20:50:37 UTC+8）

**代理总数：45**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 45 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 45 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1148ms | ✓ 1780ms | ✓ 671ms | ✓ 926ms | ✓ 750ms | http |
| 113.160.132.26:8080 | ✓ 1411ms | ✓ 1495ms | 否 | ✓ 1740ms | 否 | http |
| 106.10.55.212:1121 | ✓ 724ms | 否 | ✓ 976ms | ✓ 1155ms | 否 | http |
| 152.42.208.139:8118 | ✓ 1477ms | 否 | ✓ 1854ms | ✓ 1142ms | ✓ 863ms | http |
| 212.58.132.5:8888 | ✓ 1174ms | 否 | ✓ 1912ms | ✓ 1589ms | ✓ 1267ms | http |
| 47.85.51.197:1080 | ✓ 1894ms | 否 | ✓ 1400ms | ✓ 1402ms | ✓ 819ms | http |
| 62.113.119.14:8080 | ✓ 692ms | 否 | ✓ 682ms | ✓ 1589ms | ✓ 1220ms | http |
| 152.32.132.190:7890 | ✓ 857ms | 否 | 否 | ✓ 1576ms | ✓ 1590ms | http |
| 167.71.196.178:80 | ✓ 756ms | 否 | ✓ 744ms | ✓ 1085ms | ✓ 853ms | http |
| 34.96.238.40:8080 | ✓ 927ms | 否 | ✓ 1873ms | ✓ 1234ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1045ms | ✓ 1360ms | 否 | 否 | ✓ 1154ms | http |
| 120.92.108.86:7890 | 否 | 否 | ✓ 1748ms | ✓ 1651ms | ✓ 1550ms | http |
| 218.108.131.186:17890 | ✓ 989ms | 否 | ✓ 913ms | ✓ 1128ms | ✓ 929ms | http |
| 115.231.181.40:8128 | ✓ 888ms | 否 | 否 | ✓ 1508ms | ✓ 1029ms | http |
| 34.71.229.255:3128 | ✓ 685ms | 否 | ✓ 1541ms | ✓ 1266ms | ✓ 1326ms | http |
| 35.225.22.61:80 | ✓ 920ms | 否 | ✓ 396ms | ✓ 1184ms | 否 | http |
| 120.92.212.16:8890 | ✓ 996ms | ✓ 1261ms | 否 | 否 | ✓ 1216ms | http |
| 168.144.75.9:3128 | ✓ 1646ms | 否 | ✓ 1472ms | ✓ 1930ms | 否 | http |
| 208.87.243.199:7878 | 否 | ✓ 1425ms | ✓ 1424ms | ✓ 1319ms | 否 | http |
| 91.99.15.45:2095 | ✓ 743ms | ✓ 1988ms | ✓ 1435ms | 否 | 否 | http |
| 120.92.212.16:7890 | ✓ 1346ms | 否 | ✓ 1661ms | ✓ 1583ms | ✓ 1154ms | http |
| 84.47.150.125:1080 | ✓ 1727ms | 否 | 否 | ✓ 1632ms | ✓ 1511ms | http |
| 177.93.132.244:3128 | ✓ 1400ms | 否 | ✓ 765ms | 否 | ✓ 1774ms | http |
| 92.119.166.68:123 | ✓ 618ms | ✓ 1999ms | ✓ 1838ms | ✓ 1817ms | 否 | http |
| 38.79.118.202:33858 | 否 | ✓ 1046ms | ✓ 750ms | ✓ 815ms | ✓ 782ms | http |
| 168.110.52.228:3128 | ✓ 664ms | 否 | ✓ 1523ms | 否 | ✓ 710ms | http |
| 64.188.77.26:3128 | ✓ 1936ms | 否 | ✓ 1420ms | ✓ 1996ms | ✓ 1757ms | http |
| 221.156.27.160:3060 | 否 | 否 | ✓ 869ms | ✓ 1024ms | ✓ 856ms | http |
| 221.156.27.160:3080 | 否 | 否 | ✓ 887ms | ✓ 1016ms | ✓ 857ms | http |
| 221.156.27.160:3020 | 否 | 否 | ✓ 683ms | ✓ 1009ms | ✓ 798ms | http |
| 103.82.23.118:5198 | ✓ 1198ms | 否 | ✓ 1323ms | ✓ 1977ms | 否 | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1123ms | ✓ 1661ms | ✓ 925ms | http |
| 130.61.174.200:1080 | ✓ 1790ms | ✓ 1681ms | 否 | ✓ 1776ms | 否 | http |
| 207.254.71.62:8088 | 否 | ✓ 1935ms | 否 | ✓ 1895ms | ✓ 1681ms | http |
| 160.250.4.245:1 | ✓ 1637ms | 否 | ✓ 1446ms | ✓ 1266ms | ✓ 1035ms | http |
| 116.80.92.29:3128 | ✓ 1639ms | 否 | 否 | ✓ 1845ms | ✓ 1651ms | http |
| 20.164.75.153:8080 | ✓ 1921ms | 否 | ✓ 1344ms | 否 | ✓ 1972ms | http |
| 152.70.91.193:40000 | ✓ 1739ms | 否 | 否 | ✓ 1778ms | ✓ 1364ms | http |
| 116.63.160.98:8899 | ✓ 993ms | ✓ 1214ms | ✓ 991ms | ✓ 1301ms | 否 | http |
| 103.178.88.54:8080 | ✓ 1982ms | 否 | ✓ 1299ms | ✓ 1497ms | 否 | http |
| 61.52.131.172:8443 | ✓ 898ms | ✓ 1146ms | ✓ 1483ms | ✓ 1191ms | ✓ 945ms | http |
| 179.49.113.230:999 | 否 | ✓ 1858ms | ✓ 1566ms | ✓ 1862ms | ✓ 1385ms | http |
| 8.222.175.80:6128 | ✓ 925ms | ✓ 1743ms | 否 | ✓ 1078ms | ✓ 877ms | http |
| 223.84.151.86:30005 | 否 | 否 | ✓ 1793ms | ✓ 1612ms | ✓ 1682ms | http |
| 103.184.99.194:8080 | ✓ 1311ms | 否 | 否 | ✓ 1559ms | ✓ 1490ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
