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

最后更新：2026-04-03 08:10:26 UTC（2026-04-03 16:10:26 UTC+8）

**代理总数：43**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 43 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 43 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1360ms | ✓ 1179ms | ✓ 1368ms | ✓ 1149ms | ✓ 809ms | http |
| 208.87.243.199:7878 | ✓ 1208ms | ✓ 1632ms | ✓ 696ms | ✓ 687ms | ✓ 924ms | http |
| 1.231.81.166:3128 | ✓ 1390ms | ✓ 1911ms | ✓ 1463ms | ✓ 1033ms | ✓ 786ms | http |
| 167.103.115.102:8800 | ✓ 1422ms | 否 | ✓ 1205ms | ✓ 1055ms | ✓ 1359ms | http |
| 113.160.132.26:8080 | ✓ 1498ms | ✓ 1926ms | ✓ 1130ms | ✓ 1313ms | ✓ 995ms | http |
| 147.161.239.240:8800 | ✓ 1359ms | 否 | ✓ 864ms | ✓ 1872ms | ✓ 1594ms | http |
| 159.223.71.162:8080 | ✓ 1412ms | 否 | ✓ 1483ms | ✓ 1079ms | ✓ 1690ms | http |
| 159.223.71.162:443 | ✓ 1436ms | 否 | ✓ 1445ms | ✓ 1146ms | ✓ 1813ms | http |
| 167.103.34.108:8800 | ✓ 1459ms | 否 | ✓ 1551ms | ✓ 1492ms | ✓ 1419ms | http |
| 203.80.138.81:50000 | 否 | 否 | ✓ 1934ms | ✓ 1380ms | ✓ 1102ms | http |
| 45.167.124.52:8080 | 否 | ✓ 1986ms | ✓ 1499ms | ✓ 1892ms | 否 | http |
| 35.225.22.61:80 | ✓ 456ms | 否 | ✓ 1442ms | 否 | ✓ 965ms | http |
| 95.213.217.168:52004 | ✓ 857ms | 否 | ✓ 873ms | 否 | ✓ 1417ms | http |
| 167.103.144.127:8800 | ✓ 1166ms | 否 | ✓ 1580ms | 否 | ✓ 1589ms | http |
| 34.96.238.40:8080 | ✓ 1909ms | ✓ 1809ms | 否 | ✓ 1161ms | 否 | http |
| 45.167.125.21:999 | ✓ 1057ms | ✓ 1951ms | ✓ 1368ms | ✓ 1935ms | ✓ 1586ms | http |
| 174.140.109.250:3128 | ✓ 795ms | ✓ 1446ms | ✓ 1326ms | ✓ 1373ms | ✓ 1275ms | http |
| 180.250.219.58:53281 | 否 | 否 | ✓ 1761ms | ✓ 1973ms | ✓ 1991ms | http |
| 5.104.87.17:8051 | ✓ 1616ms | 否 | ✓ 1690ms | 否 | ✓ 1327ms | http |
| 167.103.31.122:8800 | ✓ 1534ms | 否 | ✓ 1364ms | ✓ 1735ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1065ms | 否 | ✓ 990ms | 否 | ✓ 1825ms | http |
| 168.110.52.228:3128 | 否 | 否 | ✓ 1523ms | ✓ 1792ms | ✓ 1093ms | http |
| 177.234.217.88:999 | ✓ 1936ms | 否 | ✓ 1795ms | ✓ 1994ms | ✓ 1750ms | http |
| 210.223.44.230:3128 | ✓ 1528ms | ✓ 1032ms | ✓ 1065ms | 否 | ✓ 745ms | http |
| 112.203.51.44:8082 | ✓ 1501ms | 否 | ✓ 1153ms | 否 | ✓ 1280ms | http |
| 150.249.255.91:3128 | ✓ 486ms | ✓ 1378ms | ✓ 1053ms | 否 | ✓ 1107ms | http |
| 120.92.212.16:8890 | ✓ 1651ms | 否 | ✓ 945ms | 否 | ✓ 979ms | http |
| 72.11.151.159:6005 | ✓ 404ms | 否 | ✓ 975ms | ✓ 1384ms | ✓ 1056ms | http |
| 34.101.184.164:3128 | ✓ 1203ms | 否 | ✓ 1061ms | ✓ 1409ms | ✓ 1651ms | http |
| 65.108.203.35:18080 | ✓ 975ms | 否 | 否 | ✓ 1989ms | ✓ 1426ms | http |
| 59.46.216.131:30001 | ✓ 1754ms | 否 | ✓ 1070ms | ✓ 1317ms | 否 | http |
| 115.231.181.40:8128 | ✓ 996ms | ✓ 1980ms | 否 | ✓ 1173ms | ✓ 991ms | http |
| 104.248.81.109:3128 | ✓ 1221ms | 否 | ✓ 1862ms | ✓ 1538ms | ✓ 1182ms | http |
| 45.12.151.226:2829 | ✓ 1848ms | 否 | ✓ 1377ms | 否 | ✓ 1397ms | http |
| 8.219.97.248:80 | ✓ 1834ms | 否 | ✓ 1672ms | ✓ 1348ms | 否 | http |
| 182.204.181.234:1080 | ✓ 1393ms | ✓ 1443ms | 否 | ✓ 1369ms | 否 | http |
| 178.218.105.99:3128 | ✓ 1964ms | 否 | ✓ 1817ms | 否 | ✓ 1714ms | http |
| 101.43.127.100:8877 | ✓ 1085ms | ✓ 1567ms | ✓ 1320ms | 否 | 否 | http |
| 103.39.51.190:8080 | ✓ 1806ms | 否 | ✓ 1938ms | ✓ 1760ms | ✓ 1453ms | http |
| 116.80.65.81:3172 | ✓ 1485ms | 否 | ✓ 1568ms | 否 | ✓ 1608ms | http |
| 103.76.31.2:8181 | ✓ 1735ms | 否 | 否 | ✓ 1495ms | ✓ 1470ms | http |
| 54.222.174.194:80 | 否 | ✓ 1580ms | ✓ 1406ms | ✓ 1823ms | 否 | http |
| 38.207.164.82:6005 | ✓ 1939ms | 否 | ✓ 1597ms | ✓ 1203ms | ✓ 892ms | http |

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
