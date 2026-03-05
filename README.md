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

最后更新：2026-03-05 12:38:04 UTC（2026-03-05 20:38:04 UTC+8）

**代理总数：60**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 59 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 60 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 1196ms | 否 | ✓ 752ms | 否 | ✓ 912ms | http |
| 211.171.114.154:3128 | 否 | ✓ 1374ms | ✓ 1642ms | 否 | ✓ 1164ms | http |
| 181.78.79.155:999 | 否 | ✓ 1726ms | ✓ 760ms | ✓ 1701ms | 否 | http |
| 35.225.22.61:80 | ✓ 631ms | 否 | 否 | ✓ 1073ms | ✓ 836ms | http |
| 120.92.212.16:7890 | ✓ 1891ms | 否 | 否 | ✓ 1298ms | ✓ 1274ms | http |
| 62.113.119.14:8080 | ✓ 1194ms | 否 | ✓ 1252ms | ✓ 1565ms | ✓ 1315ms | http |
| 210.223.44.230:3128 | ✓ 1761ms | ✓ 1321ms | ✓ 1148ms | ✓ 1063ms | ✓ 862ms | http |
| 47.77.193.180:1080 | ✓ 559ms | 否 | ✓ 240ms | ✓ 782ms | ✓ 592ms | http |
| 121.237.181.137:8888 | ✓ 1022ms | 否 | ✓ 1116ms | ✓ 1356ms | 否 | http |
| 120.92.212.16:8890 | 否 | ✓ 1329ms | ✓ 1041ms | ✓ 1311ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1110ms | 否 | 否 | ✓ 1816ms | ✓ 1200ms | http |
| 116.80.82.226:3172 | 否 | 否 | ✓ 1647ms | ✓ 1948ms | ✓ 1735ms | http |
| 116.80.82.220:3172 | 否 | 否 | ✓ 1645ms | ✓ 1940ms | ✓ 1760ms | http |
| 81.70.169.194:80 | ✓ 1002ms | ✓ 1415ms | ✓ 1001ms | 否 | ✓ 1010ms | http |
| 116.80.82.225:3172 | ✓ 1782ms | 否 | ✓ 1678ms | ✓ 1963ms | ✓ 1755ms | http |
| 39.104.201.40:7890 | ✓ 995ms | ✓ 1352ms | ✓ 1055ms | ✓ 1297ms | ✓ 1011ms | http |
| 14.143.222.113:10155 | ✓ 1132ms | 否 | ✓ 1109ms | ✓ 1572ms | 否 | http |
| 101.43.255.96:80 | ✓ 1421ms | 否 | ✓ 1706ms | 否 | ✓ 1073ms | http |
| 103.215.36.88:18191 | ✓ 1817ms | ✓ 1356ms | 否 | 否 | ✓ 1149ms | http |
| 150.107.140.238:3128 | ✓ 918ms | 否 | ✓ 1082ms | ✓ 1232ms | 否 | http |
| 116.80.82.232:3172 | ✓ 1578ms | 否 | ✓ 1564ms | 否 | ✓ 1742ms | http |
| 116.80.82.227:3172 | ✓ 1583ms | 否 | ✓ 1548ms | ✓ 1945ms | ✓ 1828ms | http |
| 116.80.82.230:3172 | ✓ 1572ms | 否 | ✓ 1568ms | ✓ 1932ms | ✓ 1850ms | http |
| 116.80.82.229:3172 | 否 | 否 | ✓ 1606ms | ✓ 1926ms | ✓ 1729ms | http |
| 94.176.3.43:7443 | ✓ 852ms | ✓ 1939ms | ✓ 1990ms | ✓ 1690ms | ✓ 1350ms | http |
| 193.32.178.160:57329 | ✓ 845ms | 否 | ✓ 868ms | ✓ 1185ms | 否 | http |
| 103.84.95.54:7890 | ✓ 1215ms | 否 | ✓ 713ms | ✓ 929ms | ✓ 1513ms | http |
| 165.227.5.10:8888 | ✓ 820ms | ✓ 1462ms | 否 | ✓ 844ms | ✓ 633ms | http |
| 165.232.183.96:3128 | 否 | 否 | ✓ 1809ms | ✓ 1728ms | ✓ 1417ms | http |
| 49.87.152.71:1080 | 否 | ✓ 1194ms | ✓ 1961ms | 否 | ✓ 1137ms | http |
| 35.234.17.221:8080 | ✓ 1858ms | 否 | 否 | ✓ 1669ms | ✓ 1299ms | http |
| 91.233.223.147:3128 | ✓ 1083ms | 否 | ✓ 1243ms | 否 | ✓ 1546ms | http |
| 147.75.202.36:10006 | ✓ 870ms | ✓ 1550ms | ✓ 874ms | ✓ 1579ms | ✓ 1527ms | http |
| 58.220.95.12:11743 | ✓ 1244ms | ✓ 1166ms | 否 | 否 | ✓ 1038ms | http |
| 183.237.195.130:3128 | ✓ 1811ms | 否 | 否 | ✓ 1557ms | ✓ 976ms | http |
| 5.75.196.26:40000 | ✓ 583ms | 否 | ✓ 961ms | 否 | ✓ 1967ms | http |
| 199.38.85.122:40004 | ✓ 1942ms | 否 | ✓ 1273ms | 否 | ✓ 1182ms | http |
| 74.48.78.224:2080 | ✓ 1073ms | ✓ 1964ms | ✓ 407ms | ✓ 1473ms | ✓ 956ms | http |
| 115.231.181.40:8128 | ✓ 1064ms | ✓ 1903ms | ✓ 1680ms | ✓ 1289ms | ✓ 1683ms | http |
| 45.136.198.40:3128 | ✓ 1041ms | 否 | 否 | ✓ 1735ms | ✓ 1794ms | http |
| 91.193.240.157:9877 | ✓ 1851ms | 否 | ✓ 1208ms | 否 | ✓ 1804ms | http |
| 95.85.252.153:21064 | ✓ 499ms | ✓ 1784ms | ✓ 1984ms | 否 | 否 | http |
| 121.128.121.54:3128 | ✓ 741ms | 否 | ✓ 1306ms | ✓ 1170ms | ✓ 901ms | http |
| 125.128.12.144:3128 | ✓ 796ms | 否 | ✓ 1532ms | ✓ 1131ms | ✓ 887ms | http |
| 168.235.110.63:3128 | ✓ 1013ms | 否 | ✓ 1078ms | ✓ 1180ms | ✓ 950ms | http |
| 14.56.177.44:3128 | ✓ 713ms | 否 | ✓ 1577ms | ✓ 1137ms | ✓ 940ms | http |
| 61.72.221.234:3128 | ✓ 999ms | 否 | ✓ 1387ms | 否 | ✓ 1469ms | http |
| 14.56.107.244:3128 | ✓ 1678ms | ✓ 1134ms | ✓ 1433ms | 否 | ✓ 1904ms | http |
| 61.72.221.194:3128 | ✓ 1471ms | 否 | ✓ 1201ms | ✓ 1732ms | ✓ 1479ms | http |
| 125.128.12.14:3128 | ✓ 1640ms | 否 | ✓ 1420ms | ✓ 1434ms | ✓ 1225ms | http |
| 83.219.250.8:62920 | ✓ 818ms | ✓ 1654ms | ✓ 1276ms | ✓ 1565ms | 否 | http |
| 185.191.236.162:3128 | ✓ 1682ms | 否 | 否 | ✓ 1780ms | ✓ 1151ms | http |
| 222.228.171.92:8080 | 否 | 否 | ✓ 1938ms | ✓ 1907ms | ✓ 1455ms | http |
| 172.212.68.37:3128 | ✓ 676ms | 否 | ✓ 790ms | ✓ 1374ms | ✓ 1137ms | http |
| 160.238.65.7:3128 | ✓ 1254ms | 否 | ✓ 1674ms | ✓ 1769ms | 否 | http |
| 88.80.150.82:8080 | ✓ 1496ms | ✓ 1785ms | 否 | 否 | ✓ 1758ms | https |
| 18.100.254.193:8088 | ✓ 1930ms | 否 | ✓ 1419ms | 否 | ✓ 1998ms | http |
| 103.215.36.88:16316 | ✓ 1395ms | ✓ 1357ms | ✓ 1298ms | ✓ 1569ms | ✓ 1180ms | http |
| 194.59.204.87:9080 | ✓ 1934ms | 否 | ✓ 1188ms | ✓ 1970ms | 否 | http |
| 160.250.4.13:1 | ✓ 1801ms | 否 | ✓ 1412ms | ✓ 1678ms | ✓ 1163ms | http |

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
