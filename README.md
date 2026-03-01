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

最后更新：2026-03-01 17:29:48 UTC（2026-03-02 01:29:48 UTC+8）

**代理总数：64**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 141.11.210.35:1080 | ✓ 1044ms | 否 | ✓ 1047ms | ✓ 931ms | ✓ 738ms | http |
| 148.135.85.87:1080 | ✓ 775ms | 否 | ✓ 1003ms | ✓ 1409ms | 否 | http |
| 104.238.30.45:59741 | ✓ 1762ms | 否 | ✓ 1935ms | 否 | ✓ 1999ms | http |
| 59.46.216.131:30001 | ✓ 1529ms | ✓ 1529ms | ✓ 1144ms | 否 | ✓ 1221ms | http |
| 35.225.22.61:80 | 否 | ✓ 1272ms | ✓ 132ms | ✓ 977ms | 否 | http |
| 205.209.118.30:3138 | ✓ 1455ms | 否 | ✓ 592ms | 否 | ✓ 1020ms | http |
| 196.70.95.87:3128 | ✓ 1930ms | 否 | ✓ 1914ms | ✓ 1926ms | ✓ 1579ms | http |
| 115.231.181.40:8128 | ✓ 1114ms | 否 | ✓ 1538ms | ✓ 1413ms | 否 | http |
| 190.9.109.202:999 | ✓ 1511ms | ✓ 1545ms | ✓ 1214ms | ✓ 1293ms | 否 | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1454ms | ✓ 1340ms | ✓ 1378ms | http |
| 101.43.255.96:80 | ✓ 1241ms | ✓ 1361ms | ✓ 1177ms | ✓ 1489ms | ✓ 1130ms | http |
| 81.70.169.194:80 | ✓ 1141ms | ✓ 1448ms | ✓ 1204ms | ✓ 1441ms | ✓ 1063ms | http |
| 103.104.99.29:80 | 否 | 否 | ✓ 1870ms | ✓ 1647ms | ✓ 1593ms | http |
| 103.104.99.89:80 | 否 | 否 | ✓ 1817ms | ✓ 1669ms | ✓ 1632ms | http |
| 34.101.184.164:3128 | ✓ 1841ms | 否 | ✓ 1253ms | ✓ 1619ms | ✓ 1109ms | http |
| 14.56.107.244:3128 | 否 | 否 | ✓ 1150ms | ✓ 1488ms | ✓ 1011ms | http |
| 222.228.171.92:8080 | ✓ 764ms | 否 | ✓ 1695ms | ✓ 1012ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1314ms | ✓ 1346ms | 否 | ✓ 1587ms | ✓ 1072ms | http |
| 45.125.67.37:8443 | ✓ 1645ms | 否 | 否 | ✓ 1439ms | ✓ 1323ms | http |
| 95.85.252.153:21064 | ✓ 776ms | ✓ 1871ms | ✓ 1769ms | 否 | 否 | http |
| 167.160.184.231:6005 | ✓ 218ms | 否 | ✓ 919ms | ✓ 1223ms | ✓ 958ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 942ms | ✓ 1114ms | ✓ 1004ms | http |
| 107.174.133.10:3128 | ✓ 1017ms | 否 | ✓ 1099ms | 否 | ✓ 807ms | http |
| 5.75.201.136:1080 | ✓ 551ms | ✓ 1461ms | ✓ 1831ms | ✓ 1766ms | ✓ 1625ms | http |
| 45.140.147.82:1081 | ✓ 416ms | ✓ 1822ms | ✓ 1334ms | ✓ 1974ms | ✓ 1375ms | http |
| 104.238.30.40:59741 | ✓ 1819ms | 否 | ✓ 1807ms | 否 | ✓ 1999ms | http |
| 61.72.221.124:3128 | ✓ 1927ms | ✓ 1946ms | ✓ 1726ms | 否 | 否 | http |
| 180.127.149.225:1080 | ✓ 1941ms | ✓ 1286ms | ✓ 1094ms | ✓ 1453ms | 否 | http |
| 62.113.119.14:8080 | ✓ 683ms | ✓ 1648ms | ✓ 897ms | 否 | 否 | http |
| 121.230.8.213:1080 | ✓ 1889ms | ✓ 1483ms | 否 | ✓ 1626ms | ✓ 1213ms | http |
| 121.230.8.11:1080 | ✓ 1262ms | ✓ 1710ms | ✓ 1376ms | 否 | ✓ 1136ms | http |
| 42.115.247.250:10031 | ✓ 1668ms | 否 | ✓ 1843ms | 否 | ✓ 1645ms | http |
| 107.155.65.87:13428 | ✓ 1006ms | 否 | ✓ 965ms | ✓ 1891ms | 否 | http |
| 61.72.110.94:3128 | ✓ 1372ms | ✓ 1975ms | 否 | ✓ 1436ms | ✓ 1093ms | http |
| 14.56.177.44:3128 | 否 | 否 | ✓ 1486ms | ✓ 1516ms | ✓ 1274ms | http |
| 195.123.209.48:3128 | ✓ 1109ms | 否 | 否 | ✓ 1862ms | ✓ 1817ms | http |
| 121.128.121.54:3128 | ✓ 1226ms | ✓ 1546ms | ✓ 1299ms | 否 | ✓ 1296ms | http |
| 121.128.121.184:3128 | ✓ 1894ms | 否 | ✓ 987ms | ✓ 1726ms | 否 | http |
| 104.238.30.38:59741 | ✓ 1701ms | 否 | ✓ 1779ms | 否 | ✓ 1995ms | http |
| 36.147.78.166:80 | ✓ 1764ms | ✓ 1776ms | 否 | 否 | ✓ 1520ms | http |
| 35.234.17.221:8080 | ✓ 950ms | ✓ 1442ms | 否 | ✓ 1056ms | ✓ 1266ms | http |
| 61.72.221.94:3128 | ✓ 1732ms | 否 | 否 | ✓ 1512ms | ✓ 1242ms | http |
| 104.238.30.37:59741 | ✓ 1794ms | 否 | ✓ 1938ms | 否 | ✓ 1996ms | http |
| 85.198.84.77:10808 | ✓ 1523ms | 否 | ✓ 1981ms | 否 | ✓ 1804ms | http |
| 43.165.195.107:3128 | ✓ 1614ms | 否 | ✓ 1128ms | ✓ 1351ms | ✓ 1058ms | http |
| 165.227.5.10:8888 | ✓ 214ms | ✓ 1350ms | ✓ 440ms | 否 | 否 | http |
| 61.72.110.54:3128 | ✓ 1587ms | 否 | ✓ 1341ms | ✓ 1185ms | 否 | http |
| 111.79.111.126:3128 | ✓ 1943ms | 否 | ✓ 1638ms | ✓ 1489ms | ✓ 1545ms | http |
| 85.208.108.43:2094 | ✓ 904ms | 否 | ✓ 992ms | 否 | ✓ 757ms | http |
| 36.147.78.166:443 | 否 | ✓ 1756ms | ✓ 1810ms | ✓ 1971ms | 否 | http |
| 45.136.198.40:3128 | ✓ 1135ms | 否 | ✓ 1633ms | 否 | ✓ 1882ms | http |
| 45.140.147.155:1082 | ✓ 1631ms | 否 | ✓ 507ms | ✓ 1501ms | ✓ 1213ms | http |
| 45.140.147.155:1081 | 否 | 否 | ✓ 1091ms | ✓ 1902ms | ✓ 1223ms | http |
| 103.35.188.243:3128 | 否 | ✓ 1884ms | 否 | ✓ 1032ms | ✓ 786ms | http |
| 172.212.68.37:3128 | ✓ 386ms | ✓ 1633ms | 否 | 否 | ✓ 1275ms | http |
| 103.39.51.190:8080 | ✓ 1851ms | 否 | ✓ 1392ms | ✓ 1558ms | ✓ 1489ms | http |
| 183.128.208.68:7890 | ✓ 1011ms | ✓ 1320ms | ✓ 1102ms | 否 | ✓ 1014ms | http |
| 2.56.178.131:443 | ✓ 1088ms | 否 | ✓ 1639ms | 否 | ✓ 1995ms | http |
| 192.71.213.85:9812 | ✓ 618ms | 否 | ✓ 484ms | ✓ 1572ms | 否 | http |
| 103.215.36.88:17337 | ✓ 1321ms | ✓ 1659ms | ✓ 1250ms | ✓ 1349ms | ✓ 1272ms | http |
| 203.150.113.33:8080 | 否 | 否 | ✓ 1731ms | ✓ 1833ms | ✓ 1859ms | http |
| 103.215.36.88:13763 | 否 | ✓ 1607ms | ✓ 1251ms | ✓ 1625ms | ✓ 1207ms | http |
| 177.243.209.133:999 | ✓ 660ms | 否 | ✓ 882ms | ✓ 1691ms | ✓ 1892ms | http |
| 77.83.203.5:443 | ✓ 1745ms | 否 | ✓ 1997ms | ✓ 1588ms | ✓ 1047ms | http |

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
