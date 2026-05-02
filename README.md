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

最后更新：2026-05-02 17:45:34 UTC（2026-05-03 01:45:34 UTC+8）

**代理总数：39**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 39 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 39 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1159ms | ✓ 1222ms | ✓ 1289ms | 否 | ✓ 1217ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1746ms | ✓ 1365ms | ✓ 1328ms | ✓ 1350ms | http |
| 45.167.124.71:999 | ✓ 978ms | ✓ 1948ms | ✓ 1583ms | ✓ 1764ms | ✓ 1452ms | http |
| 148.230.4.241:999 | ✓ 1730ms | 否 | ✓ 1167ms | ✓ 1764ms | 否 | http |
| 47.77.216.82:1080 | ✓ 781ms | ✓ 843ms | ✓ 502ms | 否 | 否 | http |
| 47.85.51.197:1080 | ✓ 255ms | 否 | ✓ 437ms | ✓ 1193ms | ✓ 936ms | http |
| 8.211.166.184:8081 | ✓ 1216ms | 否 | ✓ 705ms | ✓ 837ms | ✓ 679ms | http |
| 206.206.126.177:2412 | ✓ 1315ms | 否 | ✓ 733ms | ✓ 987ms | ✓ 755ms | http |
| 72.11.150.178:6005 | ✓ 828ms | ✓ 1487ms | ✓ 1077ms | 否 | ✓ 1415ms | http |
| 86.104.72.220:1081 | ✓ 453ms | ✓ 1217ms | ✓ 1213ms | ✓ 1467ms | 否 | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1063ms | ✓ 1485ms | ✓ 1208ms | http |
| 8.219.97.248:80 | ✓ 1636ms | 否 | ✓ 1779ms | ✓ 1365ms | 否 | http |
| 149.51.42.10:8080 | ✓ 818ms | ✓ 1748ms | 否 | ✓ 1497ms | 否 | http |
| 213.111.146.36:18080 | ✓ 808ms | 否 | ✓ 1654ms | ✓ 1735ms | ✓ 1414ms | http |
| 218.108.131.186:17890 | ✓ 807ms | ✓ 1023ms | ✓ 839ms | ✓ 1793ms | ✓ 839ms | http |
| 117.236.124.166:3128 | ✓ 1314ms | 否 | ✓ 1829ms | 否 | ✓ 1733ms | http |
| 103.157.200.126:3128 | ✓ 1562ms | 否 | ✓ 1437ms | 否 | ✓ 1529ms | http |
| 80.92.204.47:1081 | ✓ 1688ms | 否 | ✓ 1744ms | ✓ 1909ms | 否 | http |
| 45.140.147.82:1081 | ✓ 1112ms | 否 | ✓ 1801ms | 否 | ✓ 1955ms | http |
| 86.104.72.219:1081 | ✓ 379ms | 否 | ✓ 346ms | ✓ 1383ms | ✓ 973ms | http |
| 86.104.72.219:1082 | ✓ 374ms | 否 | ✓ 349ms | ✓ 1447ms | 否 | http |
| 62.60.231.71:56608 | ✓ 972ms | 否 | ✓ 1050ms | 否 | ✓ 1303ms | http |
| 109.120.156.122:8090 | ✓ 1718ms | 否 | ✓ 1782ms | 否 | ✓ 1774ms | http |
| 160.238.65.4:3128 | ✓ 1682ms | 否 | ✓ 1725ms | 否 | ✓ 1718ms | http |
| 86.104.74.110:1081 | ✓ 1340ms | ✓ 1375ms | ✓ 1503ms | 否 | 否 | http |
| 92.119.127.208:6005 | ✓ 751ms | 否 | ✓ 1338ms | ✓ 1986ms | 否 | http |
| 149.51.42.10:3128 | ✓ 1627ms | ✓ 1507ms | 否 | ✓ 1702ms | 否 | http |
| 152.42.177.32:8888 | ✓ 898ms | 否 | ✓ 1126ms | ✓ 1229ms | 否 | http |
| 101.32.243.189:80 | ✓ 1148ms | ✓ 1493ms | ✓ 1213ms | ✓ 1289ms | ✓ 1554ms | http |
| 160.238.65.5:3128 | ✓ 655ms | ✓ 1850ms | ✓ 1241ms | ✓ 1570ms | 否 | http |
| 121.230.8.136:1080 | ✓ 1492ms | ✓ 1502ms | ✓ 1198ms | 否 | ✓ 1337ms | http |
| 3.101.133.120:80 | ✓ 330ms | ✓ 1153ms | ✓ 731ms | ✓ 1358ms | ✓ 1284ms | http |
| 175.215.192.51:3104 | ✓ 1911ms | ✓ 1792ms | ✓ 1450ms | ✓ 1276ms | 否 | http |
| 152.32.132.190:7890 | ✓ 653ms | 否 | ✓ 1083ms | ✓ 809ms | ✓ 659ms | http |
| 210.223.44.230:3128 | 否 | ✓ 942ms | ✓ 927ms | ✓ 1709ms | ✓ 662ms | http |
| 64.188.77.26:3128 | ✓ 1513ms | 否 | ✓ 807ms | 否 | ✓ 1963ms | http |
| 59.46.216.131:30001 | ✓ 883ms | ✓ 1304ms | ✓ 1152ms | 否 | 否 | http |
| 61.52.131.172:8443 | ✓ 872ms | ✓ 1194ms | ✓ 961ms | ✓ 1205ms | ✓ 932ms | http |
| 168.110.52.228:3128 | ✓ 1949ms | 否 | 否 | ✓ 1354ms | ✓ 1273ms | http |

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
