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

最后更新：2026-05-02 18:39:01 UTC（2026-05-03 02:39:01 UTC+8）

**代理总数：47**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 47 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 47 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | 否 | ✓ 990ms | ✓ 682ms | ✓ 914ms | ✓ 1759ms | http |
| 8.211.166.184:8081 | ✓ 1452ms | 否 | ✓ 1046ms | ✓ 1026ms | ✓ 847ms | http |
| 206.206.126.177:2412 | ✓ 1556ms | ✓ 1744ms | ✓ 920ms | ✓ 1220ms | ✓ 948ms | http |
| 117.236.124.166:3128 | ✓ 1570ms | 否 | ✓ 983ms | 否 | ✓ 1185ms | http |
| 113.160.132.26:8080 | ✓ 1700ms | ✓ 1560ms | 否 | 否 | ✓ 1543ms | http |
| 45.167.124.71:999 | ✓ 1039ms | ✓ 1832ms | ✓ 1247ms | ✓ 1877ms | ✓ 1310ms | http |
| 72.11.150.178:6005 | ✓ 290ms | ✓ 1235ms | ✓ 839ms | ✓ 1325ms | ✓ 904ms | http |
| 1.231.81.166:3128 | ✓ 948ms | ✓ 1486ms | ✓ 1526ms | ✓ 1711ms | ✓ 1111ms | http |
| 20.164.75.153:8080 | ✓ 1924ms | 否 | ✓ 1867ms | 否 | ✓ 1784ms | http |
| 152.32.132.190:7890 | ✓ 1085ms | 否 | ✓ 1792ms | 否 | ✓ 886ms | http |
| 109.120.156.122:8090 | ✓ 943ms | ✓ 1988ms | ✓ 1615ms | 否 | ✓ 1537ms | http |
| 181.78.194.249:999 | ✓ 1694ms | ✓ 1863ms | ✓ 1749ms | 否 | ✓ 1678ms | http |
| 149.51.42.10:8080 | ✓ 558ms | ✓ 1424ms | 否 | ✓ 1797ms | 否 | http |
| 149.51.42.10:3128 | ✓ 648ms | ✓ 1728ms | 否 | ✓ 1276ms | 否 | http |
| 47.77.216.82:1080 | ✓ 311ms | ✓ 1099ms | ✓ 1818ms | ✓ 1063ms | ✓ 856ms | http |
| 213.111.146.36:18080 | ✓ 976ms | ✓ 1286ms | 否 | ✓ 1449ms | 否 | http |
| 218.108.131.186:17890 | ✓ 1033ms | ✓ 1515ms | ✓ 1066ms | ✓ 1393ms | ✓ 1060ms | http |
| 103.157.200.126:3128 | ✓ 977ms | 否 | ✓ 1024ms | 否 | ✓ 1761ms | http |
| 210.223.44.230:3128 | ✓ 1912ms | 否 | 否 | ✓ 1768ms | ✓ 1300ms | http |
| 80.92.204.47:1081 | ✓ 920ms | ✓ 1278ms | 否 | ✓ 1668ms | ✓ 1288ms | http |
| 103.133.254.4:3128 | 否 | 否 | ✓ 1579ms | ✓ 1932ms | ✓ 1656ms | http |
| 86.104.74.110:1081 | ✓ 1309ms | ✓ 1938ms | ✓ 1035ms | 否 | ✓ 1903ms | http |
| 212.58.132.5:8888 | ✓ 1072ms | ✓ 1936ms | ✓ 1689ms | ✓ 1488ms | ✓ 1189ms | http |
| 130.61.174.200:1080 | ✓ 965ms | 否 | ✓ 1522ms | ✓ 1609ms | 否 | http |
| 34.96.238.40:8080 | ✓ 1974ms | 否 | ✓ 1012ms | ✓ 1228ms | ✓ 1823ms | http |
| 45.38.143.90:10808 | ✓ 735ms | ✓ 1383ms | ✓ 754ms | ✓ 1682ms | ✓ 1246ms | http |
| 91.108.243.203:3128 | ✓ 686ms | ✓ 1977ms | ✓ 1107ms | 否 | ✓ 1856ms | http |
| 92.119.127.208:6005 | ✓ 1059ms | ✓ 1526ms | ✓ 1703ms | ✓ 1824ms | 否 | http |
| 133.242.16.174:3128 | ✓ 1912ms | 否 | ✓ 1689ms | ✓ 1446ms | ✓ 1202ms | http |
| 86.104.74.110:1082 | ✓ 1195ms | ✓ 1394ms | ✓ 1282ms | ✓ 1636ms | ✓ 1069ms | http |
| 3.101.133.120:80 | 否 | ✓ 1515ms | ✓ 1499ms | ✓ 1331ms | ✓ 1178ms | http |
| 38.180.2.107:3128 | ✓ 1739ms | ✓ 1641ms | ✓ 1622ms | 否 | 否 | http |
| 107.173.211.190:3128 | ✓ 1106ms | ✓ 973ms | ✓ 1283ms | ✓ 967ms | ✓ 870ms | http |
| 23.185.200.94:37514 | ✓ 1100ms | ✓ 892ms | 否 | ✓ 1040ms | ✓ 917ms | http |
| 86.104.72.219:1082 | 否 | ✓ 1889ms | ✓ 993ms | ✓ 1858ms | 否 | http |
| 92.119.127.211:6005 | ✓ 1221ms | 否 | 否 | ✓ 1988ms | ✓ 1448ms | http |
| 222.102.86.137:3028 | 否 | ✓ 1995ms | ✓ 1421ms | ✓ 1567ms | ✓ 1714ms | http |
| 86.104.72.220:1082 | ✓ 1167ms | 否 | ✓ 320ms | ✓ 1240ms | 否 | http |
| 86.104.72.220:1081 | ✓ 189ms | ✓ 1196ms | ✓ 1088ms | 否 | 否 | http |
| 62.60.231.71:56608 | ✓ 1067ms | ✓ 1881ms | ✓ 1938ms | 否 | ✓ 1129ms | http |
| 86.104.72.219:1081 | ✓ 463ms | ✓ 1729ms | ✓ 1389ms | ✓ 1446ms | ✓ 1182ms | http |
| 185.230.191.240:3128 | ✓ 881ms | ✓ 1678ms | ✓ 1883ms | ✓ 1906ms | 否 | http |
| 101.32.243.189:80 | ✓ 1377ms | 否 | ✓ 1693ms | 否 | ✓ 1510ms | http |
| 61.52.131.172:8443 | ✓ 1036ms | ✓ 1389ms | ✓ 1157ms | ✓ 1435ms | ✓ 1110ms | http |
| 62.113.119.14:8080 | ✓ 1042ms | 否 | 否 | ✓ 1646ms | ✓ 1083ms | http |
| 223.16.170.103:80 | ✓ 1724ms | 否 | ✓ 1735ms | ✓ 1801ms | 否 | http |
| 160.238.65.4:3128 | ✓ 1686ms | 否 | ✓ 1270ms | ✓ 1675ms | 否 | http |

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
